package main

import "fmt"

func main() {
	//Create a value and assign it to a variable.
	//	Print the address of that value.

	x := 42
	fmt.Println("Address is: ", &x)

}
