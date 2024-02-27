package main

import "fmt"

func intersecton[T comparable](lhs, rhs []T) []T {
	// В мапе будут храниться встреченные значения и сколько раз встречались
	m := map[T]int{}
	for _, v := range lhs {
		m[v]++
	}
	for _, v := range rhs {
		m[v]++
	}
	// Собираем результат из тех, кто встречается больше одного раза
	res := []T{}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	fmt.Println(intersecton[rune]([]rune{'a', 'b', 'c'}, []rune{'b', 'c', 'd'}))
	fmt.Println(intersecton[int]([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(intersecton[int]([]int{1, 2, 3}, []int{4, 5, 6}))
}
