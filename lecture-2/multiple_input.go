package main

import "fmt"

func sum(args ...int) int {
	result := 0
	for _, it := range args {
		result += it
	}
	return result
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	slice := []int{1, 2, 3, 4, 5}

	secondSlice := []int{6, 7, 8}

	slice = append(slice, secondSlice...)

	fmt.Println(sum(slice...))

	// fmt.Println(sum(1, slice...))
}
