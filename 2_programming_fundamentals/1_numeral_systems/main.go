package main

import "fmt"

func main() {
	//Write a program that prints a number in decimal, binary, and hex

	number := 42

	fmt.Printf("%d\n", number)
	fmt.Printf("%b\n", number)
	fmt.Printf("%o\n", number)
	fmt.Printf("%x\n", number)

	//64  32 16  8  4  2  1
	// 0  1  0   1  0  1  0

	//64    8    0
	//0     5    2

	//16     0
	// 2     a
}
