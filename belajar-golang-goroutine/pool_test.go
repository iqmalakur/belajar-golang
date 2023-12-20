package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	var group sync.WaitGroup

	pool.Put("Iqmal")
	pool.Put("Akbar")
	pool.Put("Kurnia")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Selesai")
}
