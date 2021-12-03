package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rifqi Muhammad Aziz"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string) // create channel
	defer close(channel)         // error/success it will be executed (channel will closed after finish)

	go GiveMeResponse(channel)

	data := <-channel // get data from channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	fmt.Println("Finish get data from channel")
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	fmt.Println("Capacity of channel:", cap(channel))   // capacity of buffer
	channel <- "Xenosty"                                // send data to buffer
	channel <- "Theord"                                 // send data to buffer
	channel <- "Theord"                                 // send data to buffer
	fmt.Println("Total data in channel:", len(channel)) // length or total data in buffer

	// print data in channel buffer
	fmt.Println("Data1:", <-channel)
	fmt.Println("Data2:", <-channel)

	fmt.Println("Done")
}
