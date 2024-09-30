package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	count := 0
	finished := 0
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			mu.Lock()
			if vote {
				count++
			}
			finished++
			mu.Unlock()
		}()
	}

	// Slightly improve this waiting condition by putting the main thread to sleep for 50 ms.
	for {
		mu.Lock()
		if count >= 5 || finished == 10 {
			mu.Unlock()
			break
		}
		mu.Unlock()
		time.Sleep(50 * time.Millisecond)
	}
	if count >= 5 {
		fmt.Println("recieved 5+ votes")
	} else {
		fmt.Println("lost")
	}
}

func requestVote() bool {
	return rand.Intn(2) > 0
}
