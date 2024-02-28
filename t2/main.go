package main

import (
	"fmt"
	"sync"
)

/*
 * Написать программу, которая конкурентно рассчитает значение квадратов чисел
 * взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
 */

func main() {
	arr := []int{2, 4, 6, 8, 10}
	mux := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	for _, elem := range arr {
		wg.Add(1)
		go func(a int, m *sync.RWMutex, wg *sync.WaitGroup) {
			defer wg.Done()
			m.RLock()
			fmt.Printf("%d^2 = %d\n", a, a*a)
			m.RUnlock()
		}(elem, mux, wg)
	}
	wg.Wait()
}
