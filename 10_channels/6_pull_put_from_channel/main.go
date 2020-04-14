package main

import "fmt"

func main() {
	//write a program that
	//puts 100 numbers to a channel
	//pull the numbers off the channel and print them

	ints := make(chan int)

	go func(ch chan<- int) {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}(ints)

	func(ch <-chan int) {
		for v := range ch {
			fmt.Println(v)

		}
	}(ints)

}
