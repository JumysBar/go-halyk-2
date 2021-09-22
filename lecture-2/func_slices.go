package main

import "fmt"

func funcWithArray(arg [5]string) {
	for _, str := range arg {
		fmt.Println(str)
	}
}

func funcWithSlice(arg []string) {
	for _, str := range arg {
		fmt.Println(str)
	}
}

func changeSliceValue(arg []int, value int) {
	arg[0] = value
	fmt.Println("Slice in changeSliceValue:", arg)
}

func appendSlice(arg []int) {
	arg = append(arg, 11)
	arg[0] = 20
	fmt.Println("Slice in appendSlice:", arg, "Slice len:", len(arg), "Capacity len:", cap(arg))
}

func pointerOnSlice(arg *[]int) {
	(*arg)[0] = 6
	*arg = append(*arg, 5)
}

func main() {
	array := [5]string{
		"Hello world!",
		"Hello students!",
		"Hello Kazakhstan",
		"Hello Astana",
		"Hello halyk",
	}

	funcWithArray(array)

	// Данный вызов не разрешен
	// funcWithSlice(array)

	fmt.Println()

	slice := make([]string, 0, 5)
	slice = append(slice, "1", "2", "3", "4", "5")

	funcWithSlice(slice)
	funcWithSlice(slice[:3])

	fmt.Println()

	// Данный вызов не разрешен
	// funcWithArray(slice)

	numSlice := []int{1, 2, 3, 4}

	fmt.Println("Slice:", numSlice)
	changeSliceValue(numSlice, 10)
	fmt.Println("Slice:", numSlice)
	changeSliceValue(numSlice[2:], 15)
	fmt.Println("Slice:", numSlice)

	fmt.Println()

	fmt.Println("Slice:", numSlice, "Slice len:", len(numSlice), "Capacity len:", cap(numSlice))
	appendSlice(numSlice)
	fmt.Println("Slice:", numSlice, "Slice len:", len(numSlice), "Capacity len:", cap(numSlice))

	fmt.Println()

	fmt.Println("Slice:", numSlice, "Slice len:", len(numSlice), "Capacity len:", cap(numSlice))
	appendSlice(numSlice[1:3])
	fmt.Println("Slice:", numSlice, "Slice len:", len(numSlice), "Capacity len:", cap(numSlice))

	fmt.Println()

	slice2 := []int{1, 2, 3, 4}
	fmt.Println("Slice:", slice2, "Slice len:", len(slice2), "Capacity len:", cap(slice2))
	pointerOnSlice(&slice2)
	fmt.Println("Slice:", slice2, "Slice len:", len(slice2), "Capacity len:", cap(slice2))
}
