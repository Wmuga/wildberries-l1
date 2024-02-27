package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

// Воркер. Читает из канала, пока он не будет закрыт
func reader(c chan int, workerNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range c {
		fmt.Println(workerNum, "Data:", data)
	}
}

func main() {
	loop := true
	wg := &sync.WaitGroup{}
	var num int
	for {
		fmt.Print("Введите число воркеров: ")
		_, err := fmt.Scanf("%d", &num)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Собственное ограничение воркеров. Минимум 1, максимум 20
		num = max(min(num, 20), 1)
		break
	}
	// Запуск работяг
	dataChannel := make(chan int, 1)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go reader(dataChannel, i, wg)
	}
	// Канал слушаний SIGINT
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	// Остановит отправку сообщений
	go func() {
		<-sigChannel
		loop = false
	}()
	// Пока не закрыт канал - отправляем всякого
	for loop {
		dataChannel <- rand.Intn(120)
		// Ограничение спама
		time.Sleep(time.Second / 3)
	}
	close(dataChannel)
	wg.Wait()
	fmt.Println("Done")
}

// Выбран по каналу, т.к. его уже всем передали. Можно ctx
