package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		url    string
		result string
	}{
		{"http://www.homedepot.com/findit", "findit"},
		{"http://google.com/search", "search"},
		{"http://bo.bo.com/one/two/shoe/", "shoe"},
		{"http://youtube.com/path/to/file?is=12345", "file"},
	}

	for _, test := range tests {
		base := GetPathBase(test.url)
		if base != test.result {
			t.Errorf("Nope - Check out the path package!, got: %v, want: %v.", base.url, base.result)
		}
	}
}
