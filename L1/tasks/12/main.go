package main

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

import "fmt"

func main() {
	// Пример последовательности строк
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем множество для последовательности
	set := createSet(sequence)

	// Выводим результат
	fmt.Println("Set:", set)
}

func createSet(elements []string) map[string]bool {
	set := make(map[string]bool)

	for _, element := range elements {
		set[element] = true
	}

	return set
}
