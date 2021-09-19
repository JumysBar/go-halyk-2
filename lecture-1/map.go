package main

import (
	"fmt"
)

func main() {
	var m1 map[string]int
	var m2 map[int]string
	// Какие значение у m1 и m2?

	m1 = map[string]int{
		"key1": 123,
		"key2": 321,
	}

	m2 = make(map[int]string)

	m2[1050] = "Hello world"
	m1["key3"] = 213

	fmt.Println(m1)
	fmt.Println(m2)
}
