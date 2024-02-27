package main

import "fmt"

func main() {
	a := 5
	b := 2
	fmt.Println("A:", a, "B:", b)
	a, b = b, a
	fmt.Println("A:", a, "B:", b)
}
