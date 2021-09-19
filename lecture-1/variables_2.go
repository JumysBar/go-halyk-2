package main

import (
	"fmt"
)

var x, y int

func main() {
	fmt.Println("x = ", x, "y = ", y)

	{
		x := true
		y := "string"

		fmt.Println("x = ", x, "y = ", y)
	}

	fmt.Println("x = ", x, "y = ", y)

	k, l := 3.14, 8+9i

	fmt.Println("k = ", k, "l = ", l)
}
