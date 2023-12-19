package main

// Реализовать пересечение двух неупорядоченных множеств.

import "fmt"

func main() {
	// Пример двух неупорядоченных множеств
	set1 := map[int]bool{1: true, 2: true, 3: true, 4: true}
	set2 := map[int]bool{3: true, 4: true, 5: true, 6: true}

	// Выполняем операцию пересечения
	result := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Intersection:", result)
}

func intersection(set1, set2 map[int]bool) map[int]bool {
	result := make(map[int]bool)

	// Перебираем элементы первого множества
	for key := range set1 {
		// Если элемент есть во втором множестве, добавляем его в результат
		if set2[key] {
			result[key] = true
		}
	}

	return result
}
