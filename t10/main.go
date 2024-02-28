package main

import (
	"fmt"
	"math"
)

/*
 * Дана последовательность температурных колебаний:
 * -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
 * Объединить данные значения в группы с шагом в 10 градусов.
 * Последовательность в подмножноствах не важна.
 */

func getDozens(a float64) int {
	// Нужно считать по модулю, т.к. math.Floor(-2.5) = 3
	b := math.Floor(math.Abs(a) / 10)
	doz := int(b) * 10
	if a < 0 {
		return -doz
	}
	return doz
}

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Будут объединены в мапу массивов
	united := map[int][]float64{}
	for _, elem := range arr {
		doz := getDozens(elem)
		// Если уже массив есть - добавляем
		if ar, ex := united[doz]; ex {
			united[doz] = append(ar, elem)
			continue
		}
		// Иначе создаем новый
		united[doz] = []float64{elem}
	}
	fmt.Println(united)
}
