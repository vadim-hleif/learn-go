package main

import "fmt"

func foo(a ...int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}

func bar(a []int) int {
	return foo(a...)
}

func main() {
	//create a func with the identifier foo that
	//takes in a variadic parameter of type int
	//pass in a value of type []int into your func (unfurl the []int)
	//returns the sum of all values of type int passed in
	fmt.Println(foo([]int{1, 2, 3, 4, 5}...))

	//create a func with the identifier bar that
	//takes in a parameter of type []int
	//returns the sum of all values of type int passed in

	fmt.Println(bar([]int{1, 2, 3, 4, 5}))

}
