package main

import (
	"fmt"
)

func main() {
	myMap := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}

	for key, val := range myMap {
		fmt.Printf("Map key: %s. Map value: %d\n", key, val)
	}
}
