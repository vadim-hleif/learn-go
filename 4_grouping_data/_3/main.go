package main

import "fmt"

func main() {
	//Using the code from the previous example, use SLICING to create the following new slices which are then printed:
	//[42 43 44 45 46]
	//[47 48 49 50 51]
	//[44 45 46 47 48]
	//[43 44 45 46 47]

	x := []int{42, 41, 40, 39, 38, 37, 36, 35, 34, 33}

	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
