package main

import (
	"fmt"
)

func main() {
	var x int
	if fmt.Scanf("%d", &x); x > 5 {
		fmt.Println("x greater than 5")
	} else {
		fmt.Println("x less than 5")
	}
}
