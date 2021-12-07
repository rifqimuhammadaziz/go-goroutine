package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

// func TestTicker(t *testing.T) {
// 	ticker := time.NewTicker(1 * time.Second)

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		ticker.Stop()
// 	}()

// 	// function will be run every 1 second (var ticker)
// 	for time := range ticker.C {
// 		fmt.Println(time)
// 	}
// }

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)

	// function will be run every 1 second (var ticker)
	for time := range channel {
		fmt.Println(time)
	}
}
