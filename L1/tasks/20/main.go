package main

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

import (
	"fmt"
	"strings"
)

func main() {
	originalString := "snow dog sun"
	reversedString := ReverseWords(originalString)

	fmt.Printf("Original string: %s\n", originalString)
	fmt.Printf("Reversed words: %s\n", reversedString)
}

// ReverseWords переворачивает слова в строке
func ReverseWords(input string) string {
	// Разбиваем строку на слова
	words := strings.Fields(input)

	// Переворачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Формируем новую строку из перевернутых слов
	return strings.Join(words, " ")
}
