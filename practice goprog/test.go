package main

import (
	"fmt"
	"packages/calculator"
)

func main() {

	number1 := 9
	number2 := 5

	// use the add function of calculator package
	fmt.Println(calculator.Add(number1, number2))

	// use the subtract function of calculator package
	fmt.Println(calculator.Subtract(number1, number2))

}
