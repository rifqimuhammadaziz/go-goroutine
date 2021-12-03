package gogoroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // running using goroutine (as asynchronous / if the program finishes first, then the go routine didn't have time to execute)
	fmt.Println("Done")

	time.Sleep(1 * time.Second) // waiting time (for waiting goroutine done)
}
