package main

import (
	"fmt"
	"math/rand"
)

func main() {
	exerciseOne()
	exerciseTwo()

	// Exercise 3:
	// total := 0
	// for i := 0; i < 10; i++ {
	// 	total := total + i
	// 	fmt.Println(total)
	// }
	// fmt.Println(total)
	// Explanation: the variable named total that is declared inside the for loop has a scope that only exists
	// within the body of the for loop. The variable named total that is used on the right hand side of the
	// declaration+assignment operation, `:=`, is from the main body scope. As you can see we never update it.
	// Lastly when we print total after the for loop it is the same variable from the beggining of the main
	// function and thus we never update it which is why it will print out zero.
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
