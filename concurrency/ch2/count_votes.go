package main

import (
	"fmt"
	"math/rand"
	"sync"
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

	// We are burning up a buncha CPU.
	// We are consistenly wasting CPU cycles to check updates on count+finished.
	for {
		mu.Lock()
		if count >= 5 || finished == 10 {
			mu.Unlock()
			break
		}
		mu.Unlock()
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
