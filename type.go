package main

import (
	"fmt"
	"reflect"
)

// MyMegaInteger тип для кастомизации базового типа int32
type MyMegaInteger int

// getAndIncrease возвращает текущее значение MyMegaInteger после чего инкриминирует его на единицу
func (i *MyMegaInteger) getAndIncrease() MyMegaInteger {
	retValue := *i
	*i++

	return retValue
}

// getValue возвращает текущее значение MyMegaInteger
func (i *MyMegaInteger) getValue() MyMegaInteger {
	return *i
}

func main() {
	// переменная с базовым типом int
	var basicInt int = 3

	// переменная с нашим новым типом MyMegaInteger
	var myInt MyMegaInteger = MyMegaInteger(3)

	// примеры сложения нашего типа и с базовыми
	example1 := int(myInt) + basicInt
	example2 := myInt + MyMegaInteger(basicInt)
	example3 := myInt + 3

	// результат сложения
	fmt.Println(example1, reflect.TypeOf(example1))
	fmt.Println(example2, reflect.TypeOf(example2))
	fmt.Println(example3, reflect.TypeOf(example3))

	// пример спец функций для типа int с оберткой MyMegaInteger
	fmt.Printf("%d\n", myInt.getValue())
	fmt.Printf("%d\n", myInt.getAndIncrease())
	fmt.Printf("%d\n", myInt.getValue())
}

/* output
	6 main.MyMegaInteger
	6 int
	6 main.MyMegaInteger

	3
	3
	4
*/
