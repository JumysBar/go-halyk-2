package main

import "fmt"

func dontPanic(y int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}()

	// арифметическая паника
	x := 5 / y
	fmt.Println(x)

	if x == 1 {
		// паника, генерируемая функцией
		panic("Panic at the disco!")
	}

	// паника с памятью
	var z []int
	z[0] = 1
}

func main() {
	var y int
	fmt.Scanf("%d\n", &y)
	dontPanic(y)
	fmt.Println("End of the main function")
}
