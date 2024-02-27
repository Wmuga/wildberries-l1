package main

import (
	"fmt"
	"math/rand"
)

// Паттерн Адаптер (Adapter) предназначен для преобразования интерфейса
// одного класса в интерфейс другого. Благодаря реализации данного паттерна
// мы можем использовать вместе классы с несовместимыми интерфейсами.

// Реализация - Включение уже существующего класса в другой класс.

type IGenerator interface {
	Next() int
}

// Структура, реализующая интерфейс
type GoGenerator struct{}

func (g *GoGenerator) Next() int {
	return rand.Int()
}

// Структура с другим интерфейсом
type OwnGenerator struct {
	seed int
}

func NewOwnGenerator(seed int) *OwnGenerator {
	return &OwnGenerator{seed}
}

func (g *OwnGenerator) Generate() int {
	g.seed *= g.seed
	if g.seed == 0 {
		g.seed = 1
	}
	return g.seed
}

// Стурктура - адаптер OwnGenerator к IGenerator
type OwnGeneratorAdapter struct {
	ownGen *OwnGenerator
}

func (adapter *OwnGeneratorAdapter) Next() int {
	return adapter.ownGen.Generate()
}

func NewOwnGeneratorAdapter(ownGen *OwnGenerator) *OwnGeneratorAdapter {
	return &OwnGeneratorAdapter{ownGen}
}

func main() {
	ownGen := NewOwnGenerator(123)
	var goGen IGenerator = &GoGenerator{}
	var ownGenAdapted IGenerator = NewOwnGeneratorAdapter(ownGen)
	fmt.Println(goGen.Next())
	fmt.Println(ownGenAdapted.Next())
}
