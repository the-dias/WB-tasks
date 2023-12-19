package main

// Поменять местами два числа без создания временной переменной.

import "fmt"

func main() {
	a, b := 1, 0
	fmt.Println(a, b)
	a, b = b, a
	fmt.Println(a, b)
}
