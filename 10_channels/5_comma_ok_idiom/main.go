package main

import (
	"fmt"
)

func main() {
	//Show the comma ok idiom starting with this code.
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
