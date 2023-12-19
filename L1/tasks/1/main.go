package main

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/

import "fmt"

// Struct can't be exported
type human struct {
	Name string
	Age  int
}

func (h human) Greeting() {
	fmt.Println("Hello from human")
}

// Embedding into structure
type Action struct {
	human
}

func main() {
	a := Action{}
	a.Greeting() // same as a.human.Greeting()
	a.human.Greeting()
}
