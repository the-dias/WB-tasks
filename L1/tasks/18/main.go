package main

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
	По завершению программа должна выводить итоговое значение счетчика.
*/

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment - инкрементирование счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// GetValue - получение текущего значения счетчика
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	// Создаем экземпляр счетчика
	counter := Counter{}

	// Создаем WaitGroup для ожидания завершения горутин
	var wg sync.WaitGroup

	// Количество горутин, которые будут инкрементировать счетчик
	numWorkers := 5

	// Инкрементирующие горутины
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Printf("Final counter value: %d\n", counter.GetValue())
}
