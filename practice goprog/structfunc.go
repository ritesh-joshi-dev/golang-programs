package main

import "fmt"

type Rectangle func(int, int) int

// create structure
type rectanglePara struct {
	length  int
	breadth int
	color   string
	rect    Rectangle
}

func main() {

	// assign values to struct variables
	result := rectanglePara{
		length:  10,
		breadth: 20,
		color:   "Red",
		rect: func(length int, breadth int) int {
			return length * breadth
		},
	}

	fmt.Println("Color of Rectangle: ", result.color)
	fmt.Println("Area of Rectangle: ", result.rect(result.length, result.breadth))
}
