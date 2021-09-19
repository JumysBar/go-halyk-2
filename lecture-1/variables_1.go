package main

import (
	"fmt"
)

func main() {
	// объявление переменных
	var x, y int

	fmt.Println("x = ", x, "y = ", y)

	// присваивание переменных
	x = 3
	x, y = 2, 4

	fmt.Println("x = ", x, "y = ", y)

	// инициализация переменных
	var k, l int = 5, 6

	fmt.Println("k = ", k, "l = ", l)
}
