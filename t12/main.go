package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

// Set будет состоять из мапы, значение - пустая структура
type Set[T comparable] struct {
	m map[T]struct{}
}

// Создать новый set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

// Добавить значение. Возвращает false если такое значение уже в set.
func (s *Set[T]) Append(value T) bool {
	if s.Exists(value) {
		return false
	}

	s.m[value] = struct{}{}

	return true
}

// Существует ли такое значение уже в set
func (s *Set[T]) Exists(value T) bool {
	var _, exists = s.m[value]
	return exists
}

// Удаляет значение из set. Возвращает true если такой элемент уже был
func (s *Set[T]) Remove(value T) bool {
	if s.Exists(value) {
		return false
	}

	delete(s.m, value)

	return true
}

// Получить все значения из set
func (s *Set[T]) Values() []T {
	return maps.Keys(s.m)
}

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet[string]()
	for _, elem := range arr {
		set.Append(elem)
	}
	fmt.Println(set.Values())
}
