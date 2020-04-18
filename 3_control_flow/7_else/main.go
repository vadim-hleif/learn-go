package main

import "fmt"

func main() {
	//	Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

	answer := 42

	if answer == 42 {
		fmt.Println("Good answer")
	} else if answer == 43{
		fmt.Println("Good answer too")
	} else {
		fmt.Println("Bad answer!")
	}

}
