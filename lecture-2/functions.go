package main

import (
	"fmt"
	"math"
)

// CalculateDistance - функция считает расстояние между двумя точками
func CalculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(Sqr(x2-x1) + Sqr(y2-y1))
}

// ScanPoint - функция для считывания координат одной точки
func ScanPoint() (x float64, y float64) {
	fmt.Printf("Enter the coordinates of the point, separated by a space\n")

	fmt.Scanf("%g %g\n", &x, &y)

	fmt.Printf("Point coordinates: %g %g\n", x, y)

	return
}

func main() {
	x1, y1 := ScanPoint()

	x2, y2 := ScanPoint()

	fmt.Printf("Distance between 2 points: %g\n", CalculateDistance(x1, y1, x2, y2))
}

// Sqr - функция считает квадрат числа
// порядок объявления не важен
func Sqr(x float64) float64 {
	return x * x
}
