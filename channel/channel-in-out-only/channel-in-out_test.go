package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

// sender (send to channel) / send only, cannot receive data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Rifqi Muhammad Aziz" // send data to channel
}

// receiver (receive from channel) / receive only, cannot send data
func OnlyOut(channel <-chan string) {
	data := <-channel // get data from channel
	fmt.Println(data)
	fmt.Println("Finish get data from channel")
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string) // create channel
	defer close(channel)         // error/success it will be executed (channel will closed after finish)

	go OnlyIn(channel)  // send data to channel
	go OnlyOut(channel) // receive data and print

	time.Sleep(5 * time.Second)
}
