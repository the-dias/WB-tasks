package main

// Реализовать собственную функцию sleep.

import (
	"fmt"
	"time"
)

// customSleep - собственная функция "sleep"
func customSleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func main() {
	fmt.Println("Before sleep")

	// Вызываем собственную функцию "sleep" на 3 секунды
	customSleep(3)

	fmt.Println("After sleep")
}
