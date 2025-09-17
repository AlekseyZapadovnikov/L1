package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

// Конструктор
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Distance возвращает евклидово расстояние между текущей точкой и другой.
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

// toString метод для вывода
func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.x, p.y)
}

func main() {
	a := NewPoint(1.5, 2.5)
	b := NewPoint(4.0, 6.0)

	fmt.Printf("a = %s\nb = %s\n", a, b)
	fmt.Printf("Distance between a and b = %.6f\n", a.Distance(b))
}
