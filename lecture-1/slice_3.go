package main

import (
	"fmt"
)

func main() {
	var x []bool
	// чему равно значение x?

	x = make([]bool, 3, 6)

	x = append(x, true)

	fmt.Printf("Slice: %v. Slice len: %v. Slice capacity: %v\n", x, len(x), cap(x))
	// Output: Slice: [false false false true]. Slice len: 4. Slice capacity: 6

	y := []bool{true, false, true}

	x = append(x, y...)

	fmt.Printf("Slice: %v. Slice len: %v. Slice capacity: %v\n", x, len(x), cap(x))
	// Output: Slice: [false false false true true false true]. Slice len: 7. Slice capacity: 16
}
