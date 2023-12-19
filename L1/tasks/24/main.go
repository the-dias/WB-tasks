package main

/*
	Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
	с инкапсулированными параметрами x,y и конструктором.
*/

import (
	"fmt"
	"math"
)

// Point представляет структуру точки в двумерном пространстве
type Point struct {
	x, y float64
}

// NewPoint - конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// Distance - метод для вычисления расстояния между двумя точками
func (p Point) Distance(q Point) float64 {
	dx := p.x - q.x
	dy := p.y - q.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем расстояние между ними
	distance := point1.Distance(point2)

	// Выводим результат
	fmt.Printf("Point 1: (%.2f, %.2f)\n", point1.x, point1.y)
	fmt.Printf("Point 2: (%.2f, %.2f)\n", point2.x, point2.y)
	fmt.Printf("Distance between points: %.2f\n", distance)
}
