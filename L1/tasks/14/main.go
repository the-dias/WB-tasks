package main

//	Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

import (
	"fmt"
	"reflect"
)

func main() {
	// Пример использования
	var intValue int = 42
	var stringValue string = "hello"
	var boolValue bool = true
	var channelValue chan int = make(chan int)

	// Тестирование функции getTypeInfo
	fmt.Println("Type of intValue:", getTypeInfo(intValue))
	fmt.Println("Type of stringValue:", getTypeInfo(stringValue))
	fmt.Println("Type of boolValue:", getTypeInfo(boolValue))
	fmt.Println("Type of channelValue:", getTypeInfo(channelValue))
}

func getTypeInfo(value interface{}) string {
	// Получаем тип переменной
	t := reflect.TypeOf(value)

	// Проверяем тип и возвращаем соответствующую строку
	switch t.Kind() {
	case reflect.Int:
		return "int"
	case reflect.String:
		return "string"
	case reflect.Bool:
		return "bool"
	case reflect.Chan:
		return "channel"
	default:
		return "unknown"
	}
}
