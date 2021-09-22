package main

import "fmt"

func printValue(x int) {
	fmt.Println("Value:", x)
}

func deferedFunc(x int) {
	defer printValue(x)

	defer printValue(x + 5)

	defer printValue(x * 2)

	if x >= 10 {
		return
	}

	fmt.Println("Some point")

	if x < 5 {
		return
	}

	fmt.Println("Some another point")
}

func deferWithPointer(x int) {
	y := 5
	defer func(z *int) {
		fmt.Printf("%d\n", *z)
	}(&y)

	if x >= 10 {
		return
	}

	y += 10

	if x < 5 {
		return
	}

	y *= 3
}

func main() {
	var x int
	fmt.Scanf("%d\n", &x)
	deferedFunc(x)

	deferWithPointer(x)
}
