package main

import (
	"fmt"
)

func main() {
	var array = []int{1, 2, 3}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}
