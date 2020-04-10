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
	//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
	defer fmt.Println(foo([]int{42, 0}...))

	fmt.Println(bar([]int{1, 2, 3, 4, 5}))

	fmt.Println("Not the last output")

}
