package main

import "fmt"

func main() {
	//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
	//	first name
	//	last name
	//	favorite ice cream flavors
	//	Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
	//

	type person struct {
		firstName               string
		lastName                string
		favoriteIceCreamFlavors []string
	}

	first := person{
		firstName:               "firstF",
		lastName:                "firstL",
		favoriteIceCreamFlavors: []string{"apple", "cookies"},
	}

	second := person{
		firstName:               "secondF",
		lastName:                "secondL",
		favoriteIceCreamFlavors: []string{"orange", "hazelnut"},
	}

	fmt.Println(first, second)

	for i, v := range first.favoriteIceCreamFlavors {
		fmt.Printf("number: %v, \t value: %v \n", i, v)
	}
}
