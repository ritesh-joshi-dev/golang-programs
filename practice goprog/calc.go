package main

import "fmt"

func main() {

	var num1, num2 int
	fmt.Println("Enter Two numbers")
	fmt.Scan(&num1, &num2)

	// + adds two variables
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)

	// - subtract two variables
	difference := num1 - num2
	fmt.Printf("%d - %d = %d\n", num1, num2, difference)

	// * multiply two variables
	product := num1 * num2
	fmt.Printf("%d * %d is %d\n", num1, num2, product)

	// / divide two integer variables
	quotient := num1 / num2
	fmt.Printf("%d / %d = %d\n", num1, num2, quotient)

	// % modulo-divides two variables
	remainder := num1 % num2
	fmt.Printf("%d %% %d = %d\n", num1, num2, remainder)

	// / divide two floating point variables
	result := float32(num1) / float32(num2)
	fmt.Printf(" Floaating point division of %d / %d = %g\n", num1, num2, result)
}
