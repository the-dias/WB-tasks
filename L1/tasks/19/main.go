package main

/*
	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
	Символы могут быть unicode.
*/
import (
	"fmt"
)

func main() {
	originalString := "главрыба — абырвалг"
	reversedString := ReverseString(originalString)

	fmt.Printf("Original string: %s\n", originalString)
	fmt.Printf("Reversed string: %s\n", reversedString)
}

// ReverseString переворачивает строку с учетом символов Unicode
func ReverseString(s string) string {
	// Преобразуем строку в слайс рун
	runes := []rune(s)

	// Получаем длину строки
	length := len(runes)

	// Переворачиваем слайс рун
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Формируем строку из слайса рун
	return string(runes)
}
