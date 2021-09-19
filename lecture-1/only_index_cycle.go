package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, 3}
	for i := range array {
		fmt.Println(array[i])
	}
}
