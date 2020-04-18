package main

import "fmt"

func main() {
	//Build and use an anonymous func

	result := func() int {
		fmt.Println("from anonymous func")
		return 42
	}()

	fmt.Println(result)
}
