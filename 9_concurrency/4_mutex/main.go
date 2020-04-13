package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//Fix the race condition you created in the previous exercise by using a mutex
	var mutex sync.Mutex
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Goroutines: \t", runtime.NumGoroutine())
	fmt.Println("CPUs: \t\t", runtime.NumCPU())
	fmt.Println("Result is\t", counter)
}
