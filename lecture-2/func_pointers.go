package main

import "fmt"

// Если функция в качестве аргумента принимает объект
// то создается копия того объекта, который передается при вызове функции
func FuncWithoutPointer(x int) {
	x = 5
	fmt.Printf("Value in FuncWithoutPointer: %d. Address: %p\n", x, &x)
}

// Если функция в качестве аргумента принимает указатель на объект,
// то копия объекта не создается и работа внутри функции идет именно с тем объектом
// на который указывает указатель
func FuncWithPointer(x *int) {
	*x = 5
	fmt.Printf("Value in FuncWithPointer: %d. Address: %p\n", *x, x)
}

// Функция вернет копия объекта
func FuncReturnValue() int {
	z := 1
	fmt.Printf("Value in FuncReturnValue: %d. Address: %p\n", z, &z)
	return z
}

// Функция вернет объект, созданный в данной функции
func FuncReturnPointer() *int {
	z := 1
	fmt.Printf("Value in FuncReturnPointer: %d. Address: %p\n", z, &z)
	return &z
}

func main() {
	y := 10

	fmt.Printf("Value: %d. Address: %p\n", y, &y)
	FuncWithoutPointer(y)
	fmt.Printf("Value: %d. Address: %p\n\n", y, &y)

	FuncWithPointer(&y)
	fmt.Printf("Value: %d. Address: %p\n\n", y, &y)

	var1 := FuncReturnValue()
	fmt.Printf("Value: %d. Address: %p\n\n", var1, &var1)

	var2 := FuncReturnPointer()
	fmt.Printf("Value: %d. Address: %p\n", *var2, var2)
}
