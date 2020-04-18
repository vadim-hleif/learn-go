package main

import "fmt"

func main() {
	//Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
	//	Access each value in the map. Print out the values, ranging over the slice.

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

	store := map[string]person{
		first.firstName:  first,
		second.firstName: second,
	}

	fmt.Println(store["firstF"], store["secondF"])

	fmt.Println()

	for key, person := range store {
		fmt.Printf("key: %v \n", key)
		fmt.Println("\t", person.firstName)
		fmt.Println("\t", person.lastName)
		for i, v := range person.favoriteIceCreamFlavors {
			fmt.Printf("\t number: %v, \tvalue: %v \n", i, v)
		}
	}

}
