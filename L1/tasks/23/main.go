package main

// Удалить i-ый элемент из слайса.

import "fmt"

func main() {
	// Пример слайса
	mySlice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	indexToRemove := 2
	fmt.Println("Original Slice:", mySlice)
	// Удаляем элемент из слайса
	newSlice := removeElement(mySlice, indexToRemove)

	// Выводим результат
	fmt.Println("Slice after removing element at index", indexToRemove, ":", newSlice)
}

func removeElement(slice []int, index int) []int {
	// Проверяем, находится ли индекс в пределах слайса
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Создаем новый слайс, исключая элемент по индексу
	return append(slice[:index], slice[index+1:]...)
}
