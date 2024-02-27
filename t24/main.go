package main

import (
	"fmt"
	"math"
)

type Point struct {
	// Инкапсулированы, не доступны
	x float64
	y float64
}

// Конструктор
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Вычисление расстояния
func (p *Point) Distance(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

// Получение координат
func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(2, 2)
	// Один способ - получить данные и посчитать самому
	p1x := p1.GetX()
	p1y := p1.GetY()
	p2x := p2.GetX()
	p2y := p2.GetY()

	dist1 := math.Sqrt(math.Pow(p1x-p2x, 2) + math.Pow(p1y-p2y, 2))
	// Другой способ - рассчет вынести в метод
	dist2 := p1.Distance(p2)

	fmt.Println(dist1, dist2)
}
