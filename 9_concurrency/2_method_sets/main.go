package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("My name is ", p.name)
}

type human interface {
	speak()
}

func saySomething(human human) {
	human.speak()
}

func main() {
	//This exercise will reinforce our understanding of method sets:
	//create a type person struct
	//	attach a method speak to type person using a pointer receiver
	//	*person
	//	create a type human interface
	//	to implicitly implement the interface, a human must have the speak method
	//	create func “saySomething”
	//	have it take in a human as a parameter
	//	have it call the speak method

	person := person{
		name: "name",
	}
	//	show the following in your code
	//	you CAN pass a value of type *person into saySomething
	saySomething(&person)
	person.speak()


	//	you CANNOT pass a value of type person into saySomething
	//error
	//saySomething(person)

}
