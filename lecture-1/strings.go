package main

import (
	"fmt"
)

func main() {
	const (
		str1 string = "это строка"
		x    int    = 1
	)
	str2 := `
		Это многострочная строка
	`
	var r rune = 'ы'
	fmt.Println(str1, x, str2, r)
}
