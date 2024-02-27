package main

import "fmt"

// Запукс сортировки
func quickSort(arr []int) {
	quickSortRec(arr, 0, len(arr)-1)
}

// Разбиение по Хоара. пивот - последний элемент1
func partition(arr []int, start, end int) int {
	pivot := arr[(start+end)/2]
	i := start
	j := end
	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

/*
 * Данная схема использует два индекса (один в начале массива, другой в конце),
 * которые приближаются друг к другу, пока не найдётся пара элементов, где один
 * больше опорного и расположен перед ним, а второй меньше и расположен после.
 * Эти элементы меняются местами. Обмен происходит до тех пор, пока индексы не пересекутся.
 * Алгоритм возвращает последний индекс
 */

// Сама сортировка.
func quickSortRec(arr []int, start, end int) {
	if end <= start {
		return
	}
	border := partition(arr, start, end)
	quickSortRec(arr, start, border)
	quickSortRec(arr, border+1, end)
}

func main() {
	arr := []int{6, 2, 8, 1, 0, 12, 6}
	quickSort(arr)
	fmt.Println(arr)
}
