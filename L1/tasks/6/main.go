package main

// Реализовать все возможные способы остановки выполнения горутины.

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// Method 1

	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-quit:
				fmt.Println("Горутина остановлена")
				return
			default:
				// выполняем некоторую работу
			}
		}
	}()
	// Отправляем сигнал для остановки горутины
	quit <- struct{}{}

	// Method 2

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine received stop signal.")
				return
			default:
				// Ваш код выполнения
				fmt.Println("Goroutine is running.")
				time.Sleep(time.Second)
			}
		}
	}()

	// Ждем некоторое время, а затем отменяем контекст
	time.Sleep(3 * time.Second)
	cancel()

	time.Sleep(time.Second) // Даем время для завершения горутины
	fmt.Println("Main function has finished.")

	// Method 3

	go func() {
		defer fmt.Println("Goroutine cleanup.")
		for {
			// Ваш код выполнения
			fmt.Println("Goroutine is running.")
			time.Sleep(time.Second)
			// Проверяем условие остановки
			if shouldStop() {
				fmt.Println("Goroutine received stop signal.")
				return
			}
		}
	}()

	// Ждем некоторое время, а затем просто завершаем программу
	time.Sleep(3 * time.Second)
	fmt.Println("Main function has finished.")

}

func shouldStop() bool {
	return true
}
