package gogoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New" // change <nil> or empty pool data with string 'New'
		},
	}

	pool.Put("Rifqi")
	pool.Put("Muhammad")
	pool.Put("Aziz")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Done...")
}
