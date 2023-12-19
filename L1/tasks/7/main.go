package main

// Реализовать конкурентную запись данных в map.
import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// Запись данных в map
	m.Store("key1", "value1")
	m.Store("key2", "value2")

	// Конкурентная запись данных в map
	go func() {
		m.Store("key3", "value3")
	}()

	go func() {
		m.Store("key4", "value4")
	}()

	// Чтение данных из map
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}
