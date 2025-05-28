package main

import (
	"fmt"
	"time"
)

func main() {

	// create two channels
	number := make(chan int)
	message := make(chan string)

	// function call with goroutine
	go channelNumber(number)
	go channelMessage(message)

	// selects and executes a channel
	select {

	case firstChannel := <-number:
		fmt.Println("Channel Data:", firstChannel)

	case secondChannel := <-message:
		fmt.Println("Channel Data:", secondChannel)

	default:
		fmt.Println("Wait!! Channels are not ready for execution")
	}

}

// goroutine to send integer data to channel
func channelNumber(number chan int) {
	time.Sleep(1 * time.Second)
	number <- 15
}

// goroutine to send string data to channel
func channelMessage(message chan string) {
	time.Sleep(2 * time.Second)
	message <- "Learning Go select"
}
