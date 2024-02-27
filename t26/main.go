package main

import (
	"fmt"
	"strings"
)

func testUnique(str string) bool {
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
