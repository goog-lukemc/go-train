package awesomeexample

import (
	"fmt"
)

func ExampleHello() {
	name := Hello("Alex")
	fmt.Println(name)

	// Output: Hello, 👋 Alex
}

func ExampleHelloEverybody() {
	greeted := HelloEverybody([]string{"Alex", "Luke"})
	for _, s := range greeted {
		fmt.Println(s)
	}

	// Unordered output:
	// Hello, 👋 Luke
	// Hello, 👋 Alex
}
