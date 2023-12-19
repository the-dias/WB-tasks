package main

// Реализовать паттерн «адаптер» на любом примере.

import "fmt"

// OldPrinter представляет структуру с несовместимым интерфейсом
type OldPrinter struct{}

// PrintOld - метод для структуры OldPrinter
func (op *OldPrinter) PrintOld(message string) {
	fmt.Println("Old Printer:", message)
}

// Printer представляет интерфейс, который мы хотим использовать
type Printer interface {
	Print(message string)
}

// Adapter представляет адаптер для OldPrinter
type Adapter struct {
	OldPrinter *OldPrinter
}

// Print - метод для адаптера, который реализует интерфейс Printer
func (a *Adapter) Print(message string) {
	a.OldPrinter.PrintOld(message)
}

func processPrinters(p Printer, message string) {
	p.Print(message)
}

func main() {
	oldPrinter := &OldPrinter{}
	adapter := &Adapter{OldPrinter: oldPrinter}

	processPrinters(adapter, "Hello, Adapter")
}
