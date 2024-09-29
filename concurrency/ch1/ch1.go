package main

import (
	"fmt"
	"time"
)

func main() {
	go periodic()
	time.Sleep(time.Second * 10)
}

func periodic() {
	i := 0
	for {
		fmt.Printf("tick %d\n", i)
		i++
		time.Sleep(time.Second)
	}
}
