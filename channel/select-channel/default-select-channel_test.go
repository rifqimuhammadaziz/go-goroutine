package gogoroutine

import (
	"fmt"
	"testing"
)

func TestDefaultSelectChannel(t *testing.T) {
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
		default:
			fmt.Println("Waiting Data...") // executed if the data has not been entered into the selected channel
		}
		if counter == 2 {
			break
		}
	}
}
