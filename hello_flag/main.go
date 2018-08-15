package main

import (
	"flag"
	"fmt"
	"os"
)

// set a variable to the value of the item supplied at the command line
var hito = flag.String("message", "gopher", "Who shall we say hi to?")

func main() {
	// Parse the commandline parameters
	flag.Parse()

	// Print the values to the screen
	fmt.Fprintf(os.Stdout, "Hello %s!\n", *hito)
}
