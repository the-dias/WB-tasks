package main

/*
	Разработать конвейер чисел.
	Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
	после чего данные из второго канала должны выводиться в stdout.
*/

import "fmt"

func main() {

	numbers := [...]int{2, 4, 6, 8, 10}

	numbersChan := make(chan int)
	squareChan := make(chan int)

	go func() {
		for _, el := range numbers {
			numbersChan <- el
		}
		close(numbersChan)
	}()

	go func() {
		for x := range numbersChan {
			squareChan <- x * 2
		}
		close(squareChan)
	}()

	for x := range squareChan {
		fmt.Println(x)
	}
}
