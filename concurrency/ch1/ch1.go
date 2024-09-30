package main

import (
	"sync"
	"time"
)

func main() {
	counter := 0
	mu := sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter = counter + 1
		}()
	}
	time.Sleep(5 * time.Second)
	mu.Lock()
	println(counter)
	mu.Unlock()
}
