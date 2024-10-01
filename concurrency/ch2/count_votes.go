package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// TODO: Spend more time on Sync.Cond
func main() {
	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	for i := 0; i < 10; i++ {
		go func() {
			vote := requestVote()
			mu.Lock()
			if vote {
				count++
			}
			finished++
			cond.Broadcast()
			mu.Unlock()
		}()
	}

	mu.Lock()
	for count < 5 && finished != 10 {
		cond.Wait()
	}
	if count >= 5 {
		fmt.Println("recieved 5+ votes!")
	} else {
		fmt.Println("lost.")
	}
	mu.Unlock()
}

func requestVote() bool {
	return rand.Intn(2) > 0
}
