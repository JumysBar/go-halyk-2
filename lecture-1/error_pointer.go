package main

import (
	"fmt"
)

func main() {
	var x *int
	*x = 10
	fmt.Printf("x = %p. *x = %d\n", x, *x)
}
