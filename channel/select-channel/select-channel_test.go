package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Xenosty"
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for { // unlimited loop
		select {
		case data := <-channel1:
			fmt.Println("Data from Channel 1:", data)
			counter++ // manual counter
		case data := <-channel2:
			fmt.Println("Data from Channel 2:", data)
			counter++ // manual counter
		}
		if counter == 2 {
			break
		}
	}
}
