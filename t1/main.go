package main

import "fmt"

/*
 * Дана структура Human (с произвольным набором полей и методов).
 * Реализовать встраивание методов в структуре Action
 * от родительской структуры Human (аналог наследования)
 */

// Структура "родительская"
type Human struct {
	Name string
}

// Структура "дочерняя"
type Action struct {
	/// Встраивание
	*Human
}

// Сама функция
func (h *Human) SayName() {
	fmt.Println("I'm", h.Name)
}

func main() {
	action := Action{
		&Human{
			Name: "Action",
		},
	}
	action.SayName()
}
