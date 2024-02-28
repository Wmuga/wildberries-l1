package main

import (
	"fmt"
	"strings"
)

/*
 * Разработать программу, которая проверяет, что все символы в строке уникальные
 * (true — если уникальные, false etc).
 * Функция проверки должна быть регистронезависимой.
 */

func testUnique(str string) bool {
	// Все уникальные складываются в мапу. Уже есть? Значит не уникальный
	uniq := map[rune]struct{}{}
	for _, elem := range strings.ToLower(str) {
		if _, ex := uniq[elem]; ex {
			return false
		}
		uniq[elem] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(testUnique("abcd"))
	fmt.Println(testUnique("abCdefAaf"))
	fmt.Println(testUnique("aabcd"))
}
