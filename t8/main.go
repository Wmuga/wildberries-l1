package main

import (
	"errors"
	"fmt"
)

var (
	ErrTooLargeBit = errors.New("too large bit number")
	ErrNegativeBit = errors.New("bit number is negative")
)

func SetBit(dest *int64, bit int) error {
	// Проверка, что бит в пределах [0,63]
	if bit < 0 {
		return ErrNegativeBit
	}
	if bit > 63 {
		return ErrNegativeBit
	}
	// Установка бита
	mask := int64(1) << bit
	*dest |= mask

	return nil
}

func main() {
	var num int64
	err := SetBit(&num, 62)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(num)
}
