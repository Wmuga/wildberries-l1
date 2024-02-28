package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
 * Разработать программу, которая переворачивает слова в строке.
 * Пример: «snow dog sun — sun dog snow»
 */

func reverseWordsString(in string) string {
	// представляем строку как слайс рун
	data := strings.Split(in, " ")
	// и реверсим по рунам
	for i := 0; i < len(data)/2; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
	return strings.Join(data, " ")
}

func main() {
	// Чтение полной строки
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Убрать \n
	str = strings.Trim(str, "\n")
	revStr := reverseWordsString(str)
	fmt.Println(str, "-", revStr)
}
