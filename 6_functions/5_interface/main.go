package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	r float64
}

func (figure square) area() float64 {
	return figure.length * figure.width
}

func (figure circle) area() float64 {
	return figure.r * figure.r * math.Pi
}

type shape interface {
	area() float64
}

func info(shape shape) {
	fmt.Println("Area of ", shape, " is ", shape.area())
}

func main() {
	//create a type SQUARE
	//create a type CIRCLE
	//attach a method to each that calculates AREA and returns it
	//circle area= π r 2
	//square area = L * W

	//create a type SHAPE that defines an interface as anything that has the AREA method
	//create a func INFO which takes type shape and then prints the area

	//create a value of type square
	//create a value of type circle

	square := square{
		length: 10,
		width:  2,
	}

	circle := circle{
		r: 5,
	}

	info(square)
	info(circle)
	//use func info to print the area of square
	//use func info to print the area of circle

}
