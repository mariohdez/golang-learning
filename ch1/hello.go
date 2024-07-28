package main

import "fmt"

// Within a go module you can have multiple packages.
// a go module is the source code and
// an exact specification of the dependencies of the code within the module.
// within a go module, code is organized into 1 or more packages.
// the main package contains the code that starts a Go program.
func main() {
	fmt.Printf("Hello %s\n", "world")
}
