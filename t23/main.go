package main

import (
	"fmt"
	"slices"
)

/*
 * Удалить i-ый элемент из слайса.
 */

func main() {
	i := 3
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	a = append(a[:i], a[i+1:]...) // новый показатель len
	slices.Delete(b, i, i+1)      // старый показатель len
	fmt.Println(a, b)
}
