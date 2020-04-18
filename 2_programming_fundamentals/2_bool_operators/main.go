package main

import "fmt"

func main() {
	//Using the following operators, write expressions and assign their values to variables:
	//==
	//<=
	//>=
	//!=
	//<
	//>

	const (
		a = 1 == 1
		b = 3 <= 3
		c = 3 >= 3
		d = 4 != 5
		e = 4 < 5
		f = 4 > 5
	)

	//Now print each of the variables.

	fmt.Println(a, b, c, d, e, f)

}
