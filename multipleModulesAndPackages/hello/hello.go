package main

import (
	"fmt"

	"examples.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	greetings.Cook("bacon")
}
