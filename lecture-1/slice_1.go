package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

	arr2 := arr1[1:4]

	fmt.Printf("Slice: %v\n", arr2)
}
