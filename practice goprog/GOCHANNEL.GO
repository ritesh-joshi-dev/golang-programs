package main

import "fmt"

func main() {

	number := make(chan int)
	message := make(chan string)

	go channeldata(number, message)
	//data received
	fmt.Println("channel Data:", <-number)
	fmt.Println("channel Data:", <-message)
}

func channeldata(number chan int, message chan string) {
	// data send
	number <- 25
	message <- "Learning Go lang"
	fmt.Println("Message sent! Send Operation Complete")
}
