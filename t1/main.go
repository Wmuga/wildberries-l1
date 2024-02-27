package main

import "fmt"

// Структура "родительская"
type Human struct {
	Name string
}

// Структура "дочерняя"
type Action struct {
	// "Наследование"
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
