package main

import "fmt"

func printMap(m map[string]int, fName string) {
	fmt.Println("[%s]: Begin", fName)
	for key, value := range m {
		fmt.Println("Key:", key, "value:", value)
	}
	fmt.Println("[%s]: End", fName)
	fmt.Println()
}

func changeMapValue(m map[string]int) {
	m["key1"] = 123
	m["key2"] = 321
	printMap(m, "changeMapValue")
}

func deleteMapValue(m map[string]int) {
	delete(m, "key1")
	printMap(m, "deleteMapValue")
}

func main() {
	m1 := map[string]int{
		"key1": 1,
	}
	printMap(m1, "main")
	changeMapValue(m1)
	printMap(m1, "main")
	deleteMapValue(m1)
	printMap(m1, "main")
}
