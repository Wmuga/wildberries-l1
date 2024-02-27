package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Читает из канала, пока не скажут, что хватит через контекст
func reader(ctx context.Context, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-c:
			fmt.Println("Data:", data)
		}
	}
}

func main() {
	N := 5
	wg := &sync.WaitGroup{}
	// Контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(int(time.Second)*N))
	defer cancel()
	// канал данных
	dataChannel := make(chan int, 1)
	wg.Add(1)
	go reader(ctx, dataChannel, wg)
LOOP:
	// Пока не выключено - отправляем всякого в канал
	for {
		select {
		case <-ctx.Done():
			break LOOP
		default:
			dataChannel <- rand.Intn(120)
			// Ограничение спама
			time.Sleep(time.Second / 3)
		}
	}
	wg.Wait()
	fmt.Println("Done")
}
