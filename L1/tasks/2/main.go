package main

/*
	Написать программу, которая конкурентно рассчитает значение квадратов
	чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	// Use WaitGroup for wait finish gorutines
	var wg sync.WaitGroup

	for _, number := range numbers {
		// Adding +1 gorutines, which need wait
		wg.Add(1)
		go func(n int) {
			// -1 for gorutines, defer is important
			defer wg.Done()
			fmt.Printf("%d * %d = %d\n", n, n, n*n)
		}(number)
	}

	wg.Wait()
}
