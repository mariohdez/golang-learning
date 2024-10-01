package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		time.Sleep(1 * time.Second)
		var flag bool = <-c
		fmt.Printf("flag=%v", flag)
	}()
	start := time.Now()
	c <- true // blocks until other goroutine receives
	fmt.Printf("send took %v\n", time.Since(start))
}
