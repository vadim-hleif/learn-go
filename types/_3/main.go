package main

import "fmt"

//Using the code from the previous exercise,
//At the package level scope, assign the following values to the three variables
//for x assign 42
//for y assign “James Bond”
//for z assign true

var x = 42
var y = "James Bond"
var z = true

func main() {
	//in func main
	//use fmt.Sprintf to print all of the VALUES to one single string.
	//ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
	//print out the value stored by variable “s”
	s := fmt.Sprintf("%v %v %v", x, y, z)
	fmt.Println(s)
}
