package main

import "fmt"

func main() {
	//Write a program that
	//assigns an int to a variable
	//prints that int in decimal, binary, and hex
	a := 42

	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%x\n", a)
	fmt.Println("___________")

	//shifts the bits of that int over 1 position to the left, and assigns that to a variable
	//prints that variable in decimal, binary, and hex

	a = a << 1

	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%x\n", a)

}
