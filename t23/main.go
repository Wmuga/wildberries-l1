package main

import (
	"fmt"
	"slices"
)

func main() {
	i := 3
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	// По сути пересоздаем слайс без i-го элемента
	a = append(a[:i], a[i+1:]...) // новый слайс
	slices.Delete(b, i, i+1)      // модифицирует старый слайс
	fmt.Println(a, b)
}
