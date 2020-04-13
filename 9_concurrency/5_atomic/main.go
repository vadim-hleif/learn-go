package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	//Fix the race condition you created in exercise #4 by using package atomic
	var counter int64 = 0
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Goroutines: \t", runtime.NumGoroutine())
	fmt.Println("CPUs: \t\t", runtime.NumCPU())
	fmt.Println("Result is\t", counter)
}
