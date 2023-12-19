package main

/*
	Дана последовательность чисел: 2,4,6,8,10.
	Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sumCh := make(chan int)

	for _, number := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sumCh <- n * n
		}(number)
	}

	go func() {
		wg.Wait()
		close(sumCh)
	}()

	sum := 0
	for s := range sumCh {
		sum += s
	}

	fmt.Printf("Сумма квадратов чисел равна %d\n", sum)
}
