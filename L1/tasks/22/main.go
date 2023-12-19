package main

/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	// Значения a и b больше 2^20
	a.SetString("1048577", 10)
	b.SetString("2097153", 10)

	// Сложение
	sum := big.NewInt(0).Add(a, b)
	fmt.Printf("Сумма: %s\n", sum.String())

	// Вычитание
	diff := big.NewInt(0).Sub(a, b)
	fmt.Printf("Разность: %s\n", diff.String())

	// Умножение
	product := big.NewInt(0).Mul(a, b)
	fmt.Printf("Произведение: %s\n", product.String())

	// Деление
	quotient := big.NewInt(0).Div(a, b)
	fmt.Printf("Деление: %s\n", quotient.String())
}
