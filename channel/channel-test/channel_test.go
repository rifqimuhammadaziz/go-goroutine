package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // create channel
	defer close(channel)         // error/success it will be executed (channel will closed after finish)

	// goroutine using anonymous func
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Rifqi Muhammad Aziz" // send data to channel
		fmt.Println("Finish send data to channel")
	}()

	data := <-channel // get data from channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
	fmt.Println("Finish get data from channel")
}
