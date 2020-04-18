package main

import "fmt"

func curry(x int) func() int {
	return func() int {
		return x + 42
	}
}

func main() {
	//Create a func which returns a func
	//assign the returned func to a variable
	//call the returned func



	f := curry(10)

	fmt.Println(f())
	fmt.Println(curry(42)())

}
