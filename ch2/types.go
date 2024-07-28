package main

import "fmt"

func main() {
	// Exercise 1
	// var i = 20
	// var f float32 = float32(i)
	// fmt.Printf("i: %d\tf%f\n", i, f)
	// main lesson, type conversions are explicit in go.

	// Exercise 2
	// const value = 10
	// var i int = value
	// var f float32 = value
	// fmt.Printf("value:%d\ti:%d\tf:%f\n", value, i, f)
	// main lesson, literals are untyped...
	// therefore, when you declare an immutable varaible via
	// a literal, you can assign that to a value because the
	// literal is untyped.

	// Exercise 3
	var b byte = 255
	var smallI int32 = 2_147_483_647
	var bigI uint64 = 18_446_744_073_709_551_615
	fmt.Printf("b: %b\n", b)
	fmt.Printf("smallI: %d\n", smallI)
	fmt.Printf("bigI: %d\n", bigI)

	// see what happens when you add one more than their limit
	b = b + 1
	smallI = smallI + 1
	bigI = bigI + 1
	fmt.Printf("b: %b\n", b)
	fmt.Printf("smallI: %d\n", smallI)
	fmt.Printf("bigI: %d\n", bigI)
	// turns out, go doesn't panic or throw an overflow error.
	// it works like a clock, and resets to the "next" value.
	// In this case, that is the smallest possible value the
	// type can hold.
}
