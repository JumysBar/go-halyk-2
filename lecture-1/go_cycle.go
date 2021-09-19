package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, 3}
	for i, it := range array {
		fmt.Printf("Index: %d. Value: %d\n", i, it)
	}
}
