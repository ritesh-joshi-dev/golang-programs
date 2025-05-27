package main

import "fmt"

func main() {
	var name string
	var age int

	// takes input value for name
	fmt.Print("Enter your name and age ")
	fmt.Scan(&name, &age)

	fmt.Printf("Name: %s\n Age: %d ", name, age)

}
