package main

import "fmt"

func main() {

	// switch with statement
	switch day := 6; day {

	case 1:
		fmt.Println("Monday")

	case 2:
		fmt.Println("Tuesday")

	case 3:
		fmt.Println("Wednesday")

	case 4:
		fmt.Println("Thursday")

	case 5:
		fmt.Println("Friday")

	case 6:
		fmt.Println("Saturday")
		fallthrough

	case 7:
		fmt.Println("Sunday")

	default:
		fmt.Println("Invalid Day!")
	}
}
