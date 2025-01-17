package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	alice := 10000
	bob := 10000
	var mu sync.Mutex

	total := alice + bob

	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			alice -= 1
			bob += 1
			mu.Unlock()
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()
			bob -= 1
			alice += 1
			mu.Unlock()
		}
	}()

	start := time.Now()
	for time.Since(start) < time.Second {
		mu.Lock()
		if alice+bob != total {
			fmt.Printf("observed violation, alice = %d, bob = %d, sum %d\n", alice, bob, alice+bob)
		}
		mu.Unlock()
	}
}
