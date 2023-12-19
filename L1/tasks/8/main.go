package main

/*
	Дана переменная int64.
	Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

import (
	"fmt"
)

func main() {
	var number int64
	fmt.Print("Enter an integer: ")
	fmt.Scan(&number)

	var position, bitValue int
	fmt.Print("Enter the position of the bit (0-based): ")
	fmt.Scan(&position)

	fmt.Print("Enter the bit value (0 or 1): ")
	fmt.Scan(&bitValue)

	result := setBit(number, position, bitValue)

	fmt.Printf("Result after setting bit at position %d to %d: %d\n", position, bitValue, result)
}

func setBit(value int64, position int, bitValue int) int64 {
	// Создаем маску с установленным битом в позиции position
	mask := int64(1 << uint(position))

	if bitValue == 1 {
		// Устанавливаем бит в 1
		return value | mask
	} else if bitValue == 0 {
		// Устанавливаем бит в 0
		return value &^ mask
	} else {
		// Некорректное значение bitValue
		return value
	}
}
