package main

import (
	"fmt"
	"runtime"
)

func main() {
	//	Create a program that prints out your OS and ARCH. Use the following commands to run it
	//go run
	//go build
	//go install

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
