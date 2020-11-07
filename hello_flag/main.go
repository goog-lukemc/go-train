package main

import (
	"flag"
	"fmt"
)

// set a variable to the value of the item supplied at the command line
var hito = flag.String("message", "gophercon 2020", "Who shall we say hi to?")

func main() {
	// Parse the commandline parameters
	flag.Parse()

	// Print the values to the screen
	fmt.Printf("Hello %s!\n", *hito)
}

// Bonus -
// What in the world does this * mean?
// What happens is we comment out the flag.Parse() line?

// Exercise 5 Minutes -
// Using pkg.go.dev documentation, change the default value to gopher?
