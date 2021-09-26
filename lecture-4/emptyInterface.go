package main

import "fmt"

// take from https://tour.golang.org/methods/14
func main() {
	var i interface{}
	i3 := "hello"
	i2 := 42

	// обращение к функции, которая вернет нам данные и тип
	describe(i)
	describe(i2)
	describe(i3)

	// функция, которая принимает любое количество переменных и возвращает их данные и тип итеративно
	describe2(i, i2, i3)

	// похоже на первую, но используется switch-case для работы с конкретным типом данных
	describe3(i, i2, i3)
}

// describe выводит на консоль с текстом - "(Value, Type)"
func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// describe2 функция, которая принимает любое количество переменных.
// И выводит на консоль с текстом - "(Value, Type)" каждый переданный объект
func describe2(i ...interface{}) {
	for index, element := range i {
		fmt.Printf("(%v, %T)\n", element, index)
	}
}

// describe3 функция, которая принимает любое количество переменных.
// И выводит на консоль особенный текст для каждого известного заранее типа
func describe3(i ...interface{}) {
	for _, element := range i {
		switch t := element.(type) {
		case bool:
			fmt.Printf("it's a boolean %t\n", t)
		case int:
			fmt.Printf("it's a integer %d\n", t)
		case string:
			fmt.Printf("it's a string %s\n", t)
		default:
			fmt.Printf("unexpected type %T\n", t)
		}
	}
}
