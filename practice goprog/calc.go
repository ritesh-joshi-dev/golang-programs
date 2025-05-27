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
	if num2 != 0 {
		quotient := num1 / num2
		fmt.Printf("%d / %d = %d\n", num1, num2, quotient)
	} else {
		fmt.Println(" num2 must be non zero value")
	}

	// % modulo-divides two variables
	if num2 != 0 {
		remainder := num1 % num2
		fmt.Printf("%d %% %d = %d\n", num1, num2, remainder)
	} else {
		fmt.Println(" num2 must be non zero value")
	}

	// / divide two floating point variables
	if num2 != 0 {
		result := float32(num1) / float32(num2)
		fmt.Printf("Floating point division of %d / %d = %g\n", num1, num2, result)
	} else {
		fmt.Println(" num2 must be non zero value")
	}
}
