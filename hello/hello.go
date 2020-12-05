package main

import (
	"fmt"

	"github.com/russomi-labs/golang-getting-started/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
