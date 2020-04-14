package main

import (
	"fmt"
)

func main() {
	//Starting with this code, pull the values off the channel using a for range loop
	c := gen()

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
