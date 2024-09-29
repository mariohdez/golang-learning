package main

import (
	"fmt"
	"sync"
	"time"
)

var done bool
var mu sync.Mutex

func main() {
	time.Sleep(time.Second)
	println("started")
	go periodic()
	time.Sleep(time.Second * 10)
	mu.Lock()
	done = true
	mu.Unlock()
	println("cancelled")
	time.Sleep(3 * time.Second)
}

func periodic() {
	i := 0
	for {
		fmt.Printf("tick %d\n", i)
		i++
		time.Sleep(time.Second)
		mu.Lock()
		if done {
			return
		}
		mu.Unlock()

	}
}
