package main

import "fmt"

func foo() int {
	return 42
}

func bar() string {
	return "Hello world"
}

func main() {
	//Hands on exercise
	//create a func with the identifier foo that returns an int
	//create a func with the identifier bar that returns an int and a string
	//call both funcs
	//print out their results

	fmt.Println(foo(), bar())

}
