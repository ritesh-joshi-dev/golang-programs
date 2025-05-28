package main

import "fmt"

func main() {

	// declare a struct
	type Student struct {
		name string
		age  int
		per  float32
	}

	// assign value to struct while creating an instance
	s1 := Student{"John", 25, 81.10}
	fmt.Println(s1)

	// define an instance
	var s2 Student

	// assign value to struct variables
	s2 = Student{
		name: "Sara",
		age:  29,
		per:  85.40,
	}

	fmt.Println(s2)
	fmt.Println("age of sara is :", s2.age)
	fmt.Println("percentage of John is :", s1.per)

}
