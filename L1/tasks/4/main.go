package main

/*
	Реализовать постоянную запись данных в канал (главный поток).
	Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
	Необходима возможность выбора количества воркеров при старте.

	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	var numWorkers int = 5
	fmt.Println("Enter n workers: ")
	fmt.Scan(&numWorkers)

	mainWorker := make(chan any)

	abortChan := make(chan struct{})

	// Wait (CTRL + C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		close(abortChan)
	}()

	go func() {
		for {
			select {
			case mainWorker <- 1:
			case mainWorker <- "hello":
			case mainWorker <- true:
			case <-abortChan:
				close(mainWorker)
				return
			}
		}
	}()
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, mainWorker, &wg, abortChan)

	}

	wg.Wait()
	fmt.Println("All workers have finished.")
}

func worker(id int, mainWorker <-chan any, wg *sync.WaitGroup, abortChan <-chan struct{}) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-mainWorker:
			if !ok {
				// Канал закрыт, завершаем работу
				fmt.Printf("Worker %d exiting.\n", id)
				return
			}
			fmt.Printf("Worker %d received: %d\n", id, data)
		case <-abortChan:
			// Получен сигнал завершения, завершаем работу
			fmt.Printf("Worker %d exiting due to signal.\n", id)
			return
		}
	}
}
