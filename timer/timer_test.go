package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C // wait timer channel (+5 second from var timer)
	fmt.Println(time)
}

func TestAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel // wait channel (+5 second from var channel)
	fmt.Println(time)
}

func TestAfterFunc(t *testing.T) {
	group := sync.WaitGroup{}
	group.Add(1)

	// case : for delay job
	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Timer 2:", time.Now())
		group.Done()
	})
	fmt.Println("Timer 1:", time.Now())

	group.Wait()
}
