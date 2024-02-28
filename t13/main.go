package main

import "fmt"

/*
 * Поменять местами два числа без создания временной переменной.
 */

func main() {
	a := 5
	b := 2
	fmt.Println("A:", a, "B:", b)
	a, b = b, a
	fmt.Println("A:", a, "B:", b)
}
