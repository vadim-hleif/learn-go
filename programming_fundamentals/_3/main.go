package main

import "fmt"

func main() {
	//Create TYPED and UNTYPED constants. Print the values of the constants.
	const (
		a      = 1
		b int8 = -128
	)
	fmt.Printf("%v %T \n", a, a)
	fmt.Printf("%v %T \n", b, b)
}
