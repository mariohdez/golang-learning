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
