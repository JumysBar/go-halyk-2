package main

import "fmt"

func main() {
	// реализация коллекции с уникальными значениями в Golang
	setCollection := map[string]struct{}{
		"uniqElement1": {},
		"uniqElement2": {},
		"uniqElement3": {},
	}

	// без какой-либо проверки можем быть уверенными, что дубликатов не будет
	setCollection["uniqElement4"] = struct{}{}
	setCollection["uniqElement4"] = struct{}{}

	// вывод данных SET в консоль
	fmt.Printf("%v", setCollection)
}
