package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//Using goroutines, create an incrementer program
	//have a variable to hold the incrementer value
	//launch a bunch of goroutines
	//each goroutine should
	//read the incrementer value
	//store it in a new variable
	//yield the processor with runtime.Gosched()
	//increment the new variable
	//write the value in the new variable back to the incrementer variable
	//use waitgroups to wait for all of your goroutines to finish
	//the above will create a race condition.
	//	Prove that it is a race condition by using the -race flag
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			temp := counter
			runtime.Gosched()
			temp++
			counter = temp
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Goroutines: \t", runtime.NumGoroutine())
	fmt.Println("CPUs: \t\t", runtime.NumCPU())
	fmt.Println("Result is\t", counter)
}
