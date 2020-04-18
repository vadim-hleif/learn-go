package main

import "fmt"

func main() {
	//Follow these steps:
	//start with this slice
	//x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	//append to that slice this value
	//52
	x = append(x, 52)

	//print out the slice
	fmt.Println(x)

	//in ONE STATEMENT append to that slice these values
	//53
	//54
	//55
	x = append(x, 53, 54, 55)

	//print out the slice
	fmt.Println(x)

	//append to the slice this slice
	//y := []int{56, 57, 58, 59, 60}
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	//print out the slice
	fmt.Println(x)

}
