package main

import (
	"fmt"
)

const (
	green = iota // 0
	blue  = iota // 1
	white = iota // 2
	black = iota // 3
)

func main() {
	fmt.Println(green, blue, white, black)

	const (
		circle   = 100*iota + 10 // 10
		square                   // 110
		_                        // missing value
		triangle                 // 310
	)

	fmt.Println(circle, square, triangle)
}
