package main

/*
	Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
	Объединить данные значения в группы с шагом в 10 градусов.
	Последовательность в подмножноствах не важна.


	Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

*/

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0

	groups := groupTemperatures(temperatures, step)

	// Выводим результат
	for key, values := range groups {
		fmt.Printf("%d:{%v}\n", key, values)
	}
}

func groupTemperatures(temperatures []float64, step float64) map[int][]float64 {
	groups := make(map[int][]float64)

	// Сортируем значения температур
	sort.Float64s(temperatures)

	// Группируем температуры в подмножества
	for _, temp := range temperatures {
		key := int(temp/step) * int(step) // Определяем ключ для группы
		groups[key] = append(groups[key], temp)
	}

	return groups
}
