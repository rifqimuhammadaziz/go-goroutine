package gogoroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	// test add total running goroutine (+100 goroutine)
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()
	}
	// default min goroutine running : 2
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCpu)

	// change total thread, rarely used because golang is optimal at managing threads
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	// test add total running goroutine (+100 goroutine)
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(5 * time.Second)
			group.Done()
		}()
	}
	// default min goroutine running : 2
	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

	group.Wait()
}
