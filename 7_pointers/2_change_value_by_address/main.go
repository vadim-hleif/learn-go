package main

import "fmt"

//create a person struct
//create a func called “changeMe” with a *person as a parameter

type person struct {
	first string
	last  string
}

func changeMe(person *person) {
	person.first = "another first"
	person.last = "another last"

}
func main() {
	//change the value stored at the *person address
	//important
	//to dereference a struct, use (*value).field
	//p1.first
	//(*p1).first
	//“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting
	//a field (but not a method), x.f is shorthand for (*x).f.”
	//https://golang.org/ref/spec#Selectors
	//create a value of type person

	person := person{
		first: "first",
		last:  "last",
	}

	fmt.Println(person)

	//print out the value
	//in func main
	//call “changeMe”
	//in func main
	//print out the value

	changeMe(&person)

	fmt.Println(person)
}
