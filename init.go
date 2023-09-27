package main

import (
	"fmt"

	"rsc.io/quote"

	"hellobot/greetings"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(quote.Go())

	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
