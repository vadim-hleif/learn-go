package main

import "fmt"

func main() {
	//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
	//
	//"James", "Bond", "Shaken, not stirred"
	//"Miss", "Moneypenny", "Helloooooo, James."
	//
	//Range over the records, then range over the data in each record.

	result := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for i, sl := range result {
		fmt.Println("row: \t",i)
		for j, name := range sl {
			fmt.Printf("\tindex: %v, \tvalue: %v\n", j, name)
		}
	}

}
