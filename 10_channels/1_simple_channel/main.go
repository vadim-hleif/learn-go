package main

import "fmt"

func main() {
	//	get this code working:
	//	using func literal, aka, anonymous self-executing func
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	//	a buffered channel
	c = make(chan int, 1)

	c <- 42
	fmt.Println(<-c)

}
