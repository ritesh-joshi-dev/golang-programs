package main

import "fmt"

func main() {

	var name = "John"
	var ptr *string

	// assign the memory address of name to the pointer
	ptr = &name

	fmt.Println("Value of pointer is", ptr)
	fmt.Println("Address of the variable", &name)
	fmt.Println("Value of name is", *ptr)
	fmt.Println("Value of name is ", name)
}
