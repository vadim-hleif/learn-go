package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	//write a program that
	//launches 10 goroutines
	//each goroutine adds 10 numbers to a channel
	//pull the numbers off the channel and print them
	ints := make(chan int)

	go func() {
		for v := range ints {
			fmt.Println(v)
		}
	}()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 10; j++ {
				ints <- rand.Intn(100)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(ints)

}
