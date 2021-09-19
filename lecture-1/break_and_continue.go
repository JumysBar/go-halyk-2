package main

import (
	"fmt"
)

func main() {
	var input string
LOOP:
	for {
		fmt.Scanln(&input)
		for {
			if len(input) > 10 {
				fmt.Println("Break nested loop")
				break
			}
			if len(input) < 5 {
				fmt.Println("Break external loop")
				break LOOP
			}
			if len(input) == 3 {
				fmt.Println("Continue nested loop")
				continue
			}
			if len(input) == 7 {
				fmt.Println("Continue external loop")
				continue LOOP
			}
		}
	}
}
