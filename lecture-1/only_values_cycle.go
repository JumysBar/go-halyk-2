package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, 3}
	for _, it := range array {
		fmt.Println(it)
	}
}
