package main

import "fmt"

func main() {
	// exerciseOne()
	// exerciseTwo()
	exerciseThree()
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

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func exerciseThree() {
	employeeStructLiteralWithoutNames := Employee{
		"mario",
		"hernandez",
		123,
	}
	employeeStructLiteralWithNames := Employee{
		firstName: "mario",
		lastName:  "hernandez",
		id:        123,
	}
	var employeeVarDeclaraction Employee
	employeeVarDeclaraction.firstName = "mario"
	employeeVarDeclaraction.lastName = "hernandez"
	employeeVarDeclaraction.id = 123

	fmt.Println(employeeStructLiteralWithoutNames)
	fmt.Println(employeeStructLiteralWithNames)
	fmt.Println(employeeVarDeclaraction)
}
