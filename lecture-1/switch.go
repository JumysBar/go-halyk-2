package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Scanf("%d", &input)

	switch input {
	case 0:
		fmt.Println("input equals 0")
		fallthrough
	case 1:
		fmt.Println("input equals 1")
	case 2:
		fmt.Println("input equals 2")
		break
	case 3, 4, 5, 6:
		fmt.Println("input equals 3 or 4 or 5 or 6")
	default:
		fmt.Println("input equals something else")
	}
}
