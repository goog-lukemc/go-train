package main

import (
	"flag"
	"fmt"
	"strings"
)

// This is a struct declaration
type message struct {
	text *string
}

// Add a method to your new struct -
func (m *message) UpperCaseIt() string {
	return strings.ToUpper(*m.text)
}

func main() {
	// Create your new variable
	var msg message

	// Set the commandline value to the property of our struct
	msg.text = flag.String("message", "gopher", "Say hi to!")

	// Call the method that parses the items at the commandline
	flag.Parse()

	// Print it to the screen
	fmt.Printf("Hello %s!\n", msg.UpperCaseIt())
}

// Pointers in go, what are they and how do we use them?
// What is the difference between a function and a method in Go and
// why would you choose one over the other?

// Exercise 5 Minutes:
// Using the string package documentation at pkg.go.dev change the
// above code to lowercase the variable output.
