package main

import (
	"fmt"
)

func main() {
	// Starting with this code, pull the values off the channel using a select statement
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	for {
		select {
		case <-q:
			return
		case v := <-c:
			fmt.Println(v)
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()

	return c
}
