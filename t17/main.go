package main

import (
	"fmt"
	"slices"
)

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

func binSearch(arr []int, elem int) int {
	start := 0
	end := len(arr) - 1
	for start < end {
		ind := (start + end) / 2

		if arr[ind] == elem {
			return ind
		}

		if arr[ind] < elem {
			start = ind
			continue
		}

		end = ind
	}
	return -1
}

func main() {
	arr := []int{1, 9, 6, 3, 7, 2, 12, 758, 2}
	slices.Sort(arr)
	fmt.Println(arr, len(arr))
	elem := 12
	fmt.Printf("Элемент %d находится по индексу %d\n", elem, binSearch(arr, elem))
}
