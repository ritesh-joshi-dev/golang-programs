package main

import "fmt"

func main() {

	var number1, number2 int
	var result bool

	fmt.Println("Enter two numbers")
	fmt.Scan(&number1, &number2)

	// equal to operator
	result = (number1 == number2)

	fmt.Printf("%d == %d returns %t \n", number1, number2, result)

	// not equal to operator
	result = (number1 != number2)

	fmt.Printf("%d != %d returns %t \n", number1, number2, result)

	// greater than operator
	result = (number1 > number2)

	fmt.Printf("%d > %d returns %t \n", number1, number2, result)

	// less than operator
	result = (number1 < number2)

	fmt.Printf("%d < %d returns %t \n", number1, number2, result)

}
