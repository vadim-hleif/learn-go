package main

import "fmt"

func main() {
	//Create a for loop using this syntax
	//for { }
	//Have it print out the years you have been alive.

	year := 1995
	for {
		fmt.Println(year)
		year++

		if year > 2200 {
			break
		}
	}

}
