package main

import (
	"fmt"
	"log"
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

	// := create the assings to the variable and determine the type
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

	// Add to the slice without the concern for  size
	for i := 0; i < 100; i++ {
		slc = append(slc, i)
	}
	// Display the  100 items
	fmt.Println(slc)

	// View a reference to specific items
	frac := slc[9:11]
	fmt.Println(frac)

	// Maps
	// Notice the map order is random
	mymap := make(map[int]string)
	for c := 0; c < 100; c++ {
		mymap[c] = string(c)
	}

	// Lets run range over our map to get the values
	for k, v := range mymap {
		log.Printf("ASCII Number:%v ASCII Charter:%v", k, v)
	}

	// How long does it take to get an item from a map[
	start := time.Now()
	log.Println(mymap[2])
	lat := time.Since(start)
	log.Printf("Op Took %v", lat)

}
