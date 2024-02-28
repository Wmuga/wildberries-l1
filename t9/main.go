package main

import (
	"fmt"
	"sync"
)

/*
 * Разработать конвейер чисел.
 * Даны два канала: в первый пишутся числа (x) из массива,
 * во второй — результат операции x*2,
 * после чего данные из второго канала должны выводиться в stdout.
 */

// Функция 0 отдает данные из массива в первый канал
func wrk0(outChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []int{10, 20, 30, 40, 50, 60}
	for _, el := range arr {
		outChan <- el
	}

	close(outChan)
}

// Функция 1 принимает данные из первого канала и отдает результат вычисления во второй канал
func wrk1(inChan, outChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range inChan {
		outChan <- data * data
	}
	close(outChan)
}

// Функция 2 принимает данные из канала и выводит их на экран
func wrk2(inChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range inChan {
		fmt.Println(data)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	chan1 := make(chan int)
	chan2 := make(chan int)
	go wrk0(chan1, wg)
	go wrk1(chan1, chan2, wg)
	go wrk2(chan2, wg)
	wg.Wait()
}
