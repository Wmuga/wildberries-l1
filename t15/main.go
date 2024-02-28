package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
)

/*
 * К каким негативным последствиям может привести данный фрагмент кода?
 * Как это исправить? Приведите корректный пример реализации
 */

var runes = []rune("авгдеёжзийклмнопрстуфхцчшщъыьэюя")

var justString string

// Функция генерации строки
func createHugeString(length int) string {
	buf := bytes.Buffer{}
	buf.Grow(length * 4)
	for i := 0; i < length; i++ {
		_, err := buf.WriteRune(runes[rand.Intn(len(runes))])
		if err != nil {
			log.Fatalln(err)
		}
	}
	return buf.String()
}

// Функция генерации слайса рун
func createHugeRuneSlice(length int) []rune {
	buf := make([]rune, length)
	for i := range buf {
		buf[i] = runes[rand.Intn(len(runes))]
	}
	return buf
}

// someFunc из примера
func hugeStrCut() {
	v := createHugeString(1 << 10)
	justString = v[:101]
}

// вариация с []rune
func hugeSliceCut() {
	v := createHugeRuneSlice(1 << 10)
	justString = string(v[:101])
}

// вариация с возвращением строки
func hugeStrRet() string {
	v := createHugeString(1 << 10)
	return v[:101]
}

// вариация с []rune и возвращением строки
func hugeRuneRet() string {
	v := createHugeRuneSlice(1 << 10)
	return string(v[:101])
}

// Генерить сразу нужный размер
func smallStr() {
	justString = createHugeString(101)
}

// Генерить сразу нужный размер []rune
func smallRune() {
	justString = string(createHugeRuneSlice(101))
}

func main() {
	hugeStrCut()
	fmt.Println(justString, len(justString))
	hugeSliceCut()
	fmt.Println(justString, len(justString))
	str := hugeStrRet()
	fmt.Println(str, len(str))
	str = hugeRuneRet()
	fmt.Println(str, len(str))
	// Итог: Лучше если генерить и отрезать то рунный слайс. А лучше сразу нужного размера
}

// go build -gcflags="-m" t15/main.go
