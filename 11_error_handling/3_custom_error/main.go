package main

import "fmt"

type customErr struct {
	err string
}

func (err customErr) Error() string {
	return err.err
}

func foo(err error) {
	fmt.Println(err)
}

//Create a struct “customErr” which implements the builtin.error interface.
//Create a func “foo” that has a value of type error as a parameter.
//Create a value of type “customErr” and pass it into “foo”. If you need a hint, here is one.
func main() {
	customErr := customErr{
		err: "some text",
	}

	foo(customErr)
}
