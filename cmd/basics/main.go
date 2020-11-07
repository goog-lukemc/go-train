package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

func main() {
	// Basic types
	// Declaritive
	var number int
	var numberA int32
	fmt.Println(reflect.TypeOf(number))
	fmt.Println(reflect.TypeOf(numberA))

	// := create and assigns to the variable and infers the type
	numberB := 1
	numberBB := int64(4)
	fmt.Println(reflect.TypeOf(numberB).Kind())
	fmt.Println(reflect.TypeOf(numberBB).Kind())

	// Arrays
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(arr).Kind())

	// Slices
	// make and enpty slice
	slc := make([]int, 0)
	fmt.Println(reflect.TypeOf(slc).Kind())

	// Blocks on waiting on enter key
	pressEnterKey()

	// Add to the slice without the concern for size
	for i := 0; i < 100; i++ {
		slc = append(slc, i)
	}

	// Display the 100 items
	fmt.Println(slc)

	// Blocks on waiting on enter key
	pressEnterKey()

	// View a reference to specific items
	frac := slc[9:11]

	fmt.Printf("\nSlice Fraction:%v\n", frac)

	// Blocks on waiting on enter key
	pressEnterKey()

	// Maps
	// Notice the map order is random
	mymap := make(map[int]string)
	for c := 0; c < 100; c++ {
		mymap[c] = string(c)
	}

	// Let's run range over our map to get the values
	for k, v := range mymap {
		log.Printf("Unicode Number:%v Unicode Charter:%v", k, v)
	}

	// How long does it take to get an item from a map
	start := time.Now()
	log.Println(mymap[2])
	lat := time.Since(start)
	log.Printf("Op Took %v", lat)

}

func pressEnterKey() {
	fmt.Print("Press enter to continue")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

// Exercise 10 minutes:
// Write code to print out the map above in int sorted order.

// Bonus:
// Slices are an extremely helpful part of the Go language. Have a look at
// https://github.com/golang/go/wiki/SliceTricks for some awesome coding examples
// to get the most out of go slices.

// Bonus:
// Go maps and the usage of them are common in Go as in other languages. Dave Cheney
// has a great blog on them found here. https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics
