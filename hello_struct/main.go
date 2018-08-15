package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// This is a struct - This customer type
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
	fmt.Fprintf(os.Stdout, "Hello %s!\n", msg.UpperCaseIt())
}
