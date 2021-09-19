package main

import (
	"fmt"
)

func main() {
	var x *int
	fmt.Printf("x = %v\n", x)
	// Output: x = <nil>

	y := 5
	x = &y
	fmt.Printf("x = %p. y = %d. *x = %d\n", x, y, *x)
	// Output: x = 0xc000018038. y = 5. *x = 5

	*x = 4
	fmt.Printf("x = %p. y = %d. *x = %d\n", x, y, *x)
}
