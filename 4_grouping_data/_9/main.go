package main

import "fmt"

func main() {
	//Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

	store := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	store["new_one"] = []string{"blalba", "other thing", "string"}

	for key, value := range store {
		fmt.Println("name: \t", key)
		for i, v := range value {
			fmt.Printf("\t index: %v, value: %v \n", i, v)
		}
	}
}
