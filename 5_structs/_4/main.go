package main

import "fmt"

func main() {
	//Create and use an anonymous struct.

	person := struct {
		fistName string
	}{
		fistName: "name",
	}

	fmt.Println(person)

}
