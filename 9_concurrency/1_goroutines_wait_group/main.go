package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//in addition to the main goroutine, launch two additional goroutines
	//each additional goroutine should print something out
	//use waitgroups to make sure each goroutine finishes before your program exists
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("Some output")
			fmt.Println("Goroutines: \t", runtime.NumGoroutine())
			fmt.Println("CPUs: \t\t", runtime.NumCPU())
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Goroutines: \t", runtime.NumGoroutine())
	fmt.Println("CPUs: \t\t", runtime.NumCPU())
}
