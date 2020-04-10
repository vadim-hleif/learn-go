package main

import "fmt"

func curry(x int) func() int {
	return func() int {
		return x + 42
	}
}

func main() {
	//A “callback” is when we pass a func into a func as an argument. For this exercise,
	//	pass a func into a func as an argument


	f := curry(10)

	fmt.Println(f())
	fmt.Println(curry(42)())

}
