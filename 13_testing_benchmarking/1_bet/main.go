package main

import (
	"fmt"
	"learn-go/13_testing_benchmarking/1_bet/dog"
)

type canine struct {
	name string
	age  int
}

//Start with this code. Get the code ready to BET on the code (add benchmarks, examples, tests).
//Run the following in this order:
//tests
//benchmarks
//coverage
//coverage shown in web browser
//examples shown in documentation in a web browser

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}
