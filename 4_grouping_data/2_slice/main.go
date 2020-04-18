package main

import "fmt"

func main() {
	//Using a COMPOSITE LITERAL:
	//create a SLICE of TYPE int
	//assign 10 VALUES
	//Range over the slice and print the values out.
	//	Using format printing
	//print out the TYPE of the slice

	x := []int{42, 41, 40, 39, 38, 37, 36, 35, 34, 33}

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Printf("%T", x)
	fmt.Println(len(x))

}
