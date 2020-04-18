package main

import "fmt"

func sumAndThen(i ...int) func(func(int)) {
	result := 0
	for _, v := range i {
		result += v
	}

	return func(cb func(result int)) {
		cb(result)
	}
}

func main() {
	//Closure is when we have “enclosed” the scope of a variable in some code block.
	//For this hands-on exercise, create a func which “encloses” the scope of a variable:
	cb := func(result int) {
		fmt.Println(result)
	}

	result := sumAndThen(1, 2, 3, 4, 5)

	result(cb)
	result(func(result int) {
		fmt.Println("Result is: ", result)
	})
}
