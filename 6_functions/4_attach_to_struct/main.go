package main

import "fmt"

type person struct {
	first string
	last  string
	age   int8
}

func (p person) speak() {
	fmt.Println("My name is ", p.first, " and my age is ", p.age)
}

func main() {
	//Create a user defined struct with
	//	the identifier “person”
	//	the fields:
	//	first
	//	last
	//	age
	//	attach a method to type person with
	//	the identifier “speak”
	//	the method should have the person say their name and age
	//	create a value of type person
	//	call the method from the value of type person

	person := person{
		first: "name",
		last:  "lastname",
		age:   25,
	}

	person.speak()
}
