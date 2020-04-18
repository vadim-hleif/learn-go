package main

import "fmt"

func main() {
	//Assign a func to a variable, then call that func
	f := func() {
		fmt.Println("Hello from f")
	}

	f()
}
