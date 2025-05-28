package main

import "fmt"

// function definition
func calculate(n1 int, n2 int) (int, int, int) {
	sum := n1 + n2
	difference := n1 - n2
	multiplication := n1 * n2

	// return two values
	return sum, difference, multiplication
}

func main() {

	// function call
	sum, difference, multiplication := calculate(21, 13)

	fmt.Println("Sum:", sum, "Difference:", difference, "Multiplication:", multiplication)

}
