package main

import (
	"testing"
)

// TestSum_basic is a basic test case
func TestSum_basic(t *testing.T) {

	// setups args that you need for your test case
	arg := struct {
		x, y int
	}{
		1, 1,
	}

	// call our function and store the result as "got"
	got := Sum(arg.x, arg.y)

	// the expected result named as "want"
	want := 2

	// simple assertion
	if got != want {
		t.Errorf("Sum(%d, %d) = %d; want %d", arg.x, arg.y, got, want)
		// error statement recommended format
		// t.Errorf("FunctionCall(%q) = %s; want %s",arg,got,want)
	}

}

// TestSum_table showcases how you can table driven test to create a hierarchical test structure
func TestSum_table(t *testing.T) {
	// setups args that you need for your test case
	type arg struct {
		x, y int
	}

	tests := []struct {
		name string
		arg  arg
		want int
	}{
		// add all our test cases to this slice of tests, a nice structure is "name" to give your test case a descriptive name (useful for describing bug fix test cases)
		// then the "arg" that you will be passing into your test case
		// and what is the expected "want" of your test cases
		{name: "super simple math", arg: arg{x: 1, y: 1}, want: 2},
		{name: "a bit more advanced math", arg: arg{x: 2, y: 2}, want: 4},
		{name: "some really intense math", arg: arg{x: 2, y: 25}, want: 27},
		{name: "insane levels of math", arg: arg{x: 100, y: 10}, want: 110000},
	}
	for _, tt := range tests {
		// our sub test name will be what we provided above
		t.Run(tt.name, func(t *testing.T) {
			// same test assertion
			if got := Sum(tt.arg.x, tt.arg.y); got != tt.want {
				// t.Error will mark the test as failed but continue execution of further test cases
				t.Errorf("Sum(%d, %d) = %d; want %d", tt.arg.x, tt.arg.y, got, tt.want)
			}
		})
	}

}
