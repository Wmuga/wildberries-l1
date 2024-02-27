package main

import (
	"fmt"
	"os"
)

func reverseString(in string) string {
	// представляем строку как слайс рун
	data := []rune(in)
	// и реверсим по рунам
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
	return string(data)
}

func main() {
	var str string
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	revStr := reverseString(str)
	fmt.Println(str, "-", revStr)
}
