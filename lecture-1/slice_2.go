package main

import (
	"fmt"
)

func main() {
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

	arr2 := arr1[:]

	arr2[0] = 6

	fmt.Printf("Array: %v\n", arr1)
	// Output: Array: [6 2 3 4 5]

	fmt.Printf("Slice: %v. Slice len: %v. Slice capacity: %v\n", arr2, len(arr2), cap(arr2))
	// Output: Slice: [6 2 3 4 5]. Slice len: 5. Slice capacity: 5

	arr2 = append(arr2, 10)

	arr2[0] = 7

	fmt.Printf("Array: %v\n", arr1)
	// Output: Array: [6 2 3 4 5]

	fmt.Printf("Slice: %v. Slice len: %v. Slice capacity: %v\n", arr2, len(arr2), cap(arr2))
	// Output: Slice: [7 2 3 4 5 10]. Slice len: 6. Slice capacity: 10
}
