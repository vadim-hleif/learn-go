package main

import "fmt"

func main() {
	//	Create a program that uses a switch statement with the switch expression
	//	specified as a variable of TYPE string with the IDENTIFIER “favSport”.

	favSport := "some string"

	switch favSport {
	case "some string":
		fmt.Println("print")
	case "other":
		fmt.Println("not print")
	}

}
