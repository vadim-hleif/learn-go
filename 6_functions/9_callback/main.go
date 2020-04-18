package main

import "fmt"

func sumAndApply(cb func(result int), i ...int) {
	result := 0
	for _, v := range i {
		result += v
	}

	cb(result)
}

func main() {
	//A “callback” is when we pass a func into a func as an argument. For this exercise,
	//	pass a func into a func as an argument
	cb := func(result int) {
		fmt.Println(result)
	}

	sumAndApply(cb, 1, 2, 3, 4, 5)
}
