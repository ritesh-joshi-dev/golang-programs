package main

import "fmt"

func main() {

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go receiveData(ch)

	// send data to the channel
	ch <- "Data Received. Receive Operation Successful"

	fmt.Println("No data. Receive Operation Blocked")
}

func receiveData(ch chan string) {

	// receive data from the channel
	fmt.Println(<-ch)

}
