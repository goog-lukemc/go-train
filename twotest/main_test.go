package main

import (
	"errors"
	"testing"
)

func TestSum(t *testing.T) {
	tests := []struct {
		url    string
		result string
		err    error
	}{
		{"http://www.homedepot.com/findit", "findit", nil},
		{"http://google.com/search", "search", nil},
		{"http://bo.bo.com/one/two/shoe/", "shoe", nil},
		{"http://youtube.com/path/to/file?is=12345", "file", nil},
		{"", "", errors.New("parse : empty url")},
	}

	for _, test := range tests {
		// Run our function
		base, err := GetPathBase(test.url)

		// If we got an error we weren't expecting fail the test
		if err != nil {
			// Check to see if we got an error we expected
			if err.Error() != test.err.Error() {
				t.Fatalf("Expecting %v  got %v", test.err.Error(), err.Error())

			}
			t.SkipNow()
		}

		// if we got a result we expected
		if *base != test.result {
			t.Errorf("Nope - Check out the path package!, got: %v, want: %v.", test.url, test.result)
		}
	}
}
