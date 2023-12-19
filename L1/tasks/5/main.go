package main

/*
	Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
	По истечению N секунд программа должна завершаться.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Println("Enter n seconds to exit program")
	fmt.Scan(&n)

	dataChan := make(chan int)
	doneChan := make(chan bool)

	// Gorutine for writing to channel
	go func() {

		for {
			select {
			case <-doneChan:
				return
			default:
				dataChan <- rand.Intn(10)

			}
		}
	}()

	// Gorutine for reading from channel
	go func() {
		for {
			select {
			case <-doneChan:
				return
			case data := <-dataChan:
				fmt.Println(data)
			}
		}
	}()

	// Завершаем программу через N секунд
	time.Sleep(time.Duration(n) * time.Second)
	close(doneChan)

}
