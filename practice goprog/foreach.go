package main

import "fmt"

func main() {

	numbers := [5]int{11, 22, 33, 44, 55}

	for i, item := range numbers {
		fmt.Println(i, item)
	}

}
