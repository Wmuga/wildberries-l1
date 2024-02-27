package main

import (
	"fmt"
	"sync"
)

/*
 * Функция параллельного рекурсивного расчета суммы квадратов массива.
 * arr - сам массив
 * start, end - границы
 * m - мьютекс. чтобы блокировать на чтение
 * wg - рабочая группа. Нужна чтобы сообщать когда закончит работу
 * res - куда происходит запись результата. Так как вычисления параллельные, лучше запись через указатель
 */
func calcRec(arr []int, start, end int, m *sync.RWMutex, wg *sync.WaitGroup, res *int) {
	// Не забываем сообщать, что закончили при выходе
	defer wg.Done()
	// Может стоит выдавать ошибку
	if start < 0 || start >= len(arr) || end < 0 || end >= len(arr) {
		return
	}
	// Один элемент. Возвращаем квадрат
	if (end - start) == 0 {
		m.RLock()
		defer m.RUnlock()
		*res = arr[start] * arr[start]
		return
	}
	// Запуск подсчета в 2х подмассивах
	var l, r int
	wg1 := &sync.WaitGroup{}
	wg1.Add(2)
	go calcRec(arr, start, (end+start)/2, m, wg1, &l)
	go calcRec(arr, (end+start)/2+1, end, m, wg1, &r)
	wg1.Wait()
	// Запись результата
	*res = l + r
}

// Функция запуска рекурсивного подсчета
func startCalc(arr []int) int {
	mux := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	var l, r int
	wg.Add(2)
	go calcRec(arr, 0, (len(arr)-1)/2, mux, wg, &l)
	go calcRec(arr, (len(arr)-1)/2+1, len(arr)-1, mux, wg, &r)
	wg.Wait()
	return l + r
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	fmt.Println(startCalc(arr))
}
