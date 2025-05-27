package main

import "fmt"

func main() {

	// declare slice variable of type integer
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	numbers = append([]int{0}, numbers...)
	// print the slice
	fmt.Println("Numbers: ", numbers)
}
