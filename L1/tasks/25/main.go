package main

// Реализовать собственную функцию sleep.

// #include <unistd.h>
// #include <synchapi.h>
import "C"

import (
	"fmt"
	"time"
)

// customSleep - собственная функция "sleep"
func customSleep(sec int) {
	<-time.After(time.Duration(sec) * time.Second)
}

func customSleepUnix(sec uint) {
	// C.Sleep()
	C.sleep(C.uint(sec))

}

func customSleepWindows(ms uint) {
	C.Sleep(C.ulong(ms * 1000))
}

func main() {
	fmt.Println("Before sleep")

	// Вызываем собственную функцию "sleep" на 3 секунды
	customSleepUnix(10)

	fmt.Println("After sleep")
}
