package main

import (
	"fmt"
	"learn-go/12_documentation/1_doc/dog"
)

//Create a dog package. The dog package should have an exported func
//“Years” which takes human years and turns them into dog years (1 human year = 7 dog years).
//Document  your code with comments. Use this code in func main.
//run your program and make sure it works
//run a local server with godoc and look at your documentation.

func main() {
	dogYears := dog.Years(3)

	fmt.Println(dogYears)
}
