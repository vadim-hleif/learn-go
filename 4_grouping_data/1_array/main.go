package main

import "fmt"

func main() {
	//Using a COMPOSITE LITERAL:
	//create an ARRAY which holds 5 VALUES of TYPE int
	//assign VALUES to each index position.
	//	Range over the array and print the values out.
	//	Using format printing
	//print out the TYPE of the array

	x := [5]int{42, 41, 40, 39, 38}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
