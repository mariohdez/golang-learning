package main

import (
	"fmt"
	"math/rand"
)

func main() {
	exerciseOne()
	exerciseTwo()
}

func exerciseOne() {
	nums := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(101))
	}
	fmt.Println(nums)
}

func exerciseTwo() {
	nums := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(101))
	}

	for _, num := range nums {
		if num%2 == 0 && num%3 == 0 {
			fmt.Printf("[%d]: Six!\n", num)
		} else if num%2 == 0 {
			fmt.Printf("[%d]: Two!\n", num)
		} else if num%3 == 0 {
			fmt.Printf("[%d]: Three!\n", num)
		} else {
			fmt.Println("Never mind")
		}
	}
}
