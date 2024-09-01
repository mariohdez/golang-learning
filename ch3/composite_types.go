package main

import "fmt"

func main() {
	exerciseOne()
	exerciseTwo()
}

func exerciseOne() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	slice1 := greetings[:2]
	slice2 := greetings[1:4]
	slice3 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}

func exerciseTwo() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}
