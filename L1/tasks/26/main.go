package main

/*
	Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
	Функция проверки должна быть регистронезависимой.

	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

import (
	"fmt"
	"strings"
)

func main() {
	testString1 := "abcd"
	testString2 := "abCdefAaf"
	testString3 := "aabcd"

	fmt.Printf("Are all characters in \"%s\" unique? %t\n", testString1, areAllCharactersUnique(testString1))
	fmt.Printf("Are all characters in \"%s\" unique? %t\n", testString2, areAllCharactersUnique(testString2))
	fmt.Printf("Are all characters in \"%s\" unique? %t\n", testString3, areAllCharactersUnique(testString3))
}

// areAllCharactersUnique проверяет, что все символы в строке уникальны
func areAllCharactersUnique(s string) bool {
	seen := make(map[rune]struct{})

	for _, char := range strings.ToLower(s) {
		if _, exists := seen[char]; exists {
			return false
		}
		seen[char] = struct{}{}
	}

	return true
}
