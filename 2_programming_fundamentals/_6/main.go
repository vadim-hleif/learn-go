package main

import "fmt"

func main() {
	//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

	const (
		_ = iota + 2020
		year2021
		year2022
		year2023
		year2024
	)

	fmt.Println(year2021, year2022, year2023, year2024)
}
