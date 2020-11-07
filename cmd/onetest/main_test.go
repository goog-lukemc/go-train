package main

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{2, 2, 4},
		{2, 25, 27},
		{100, 10, 110},
	}

	for _, test := range tests {
		total := Sum(test.x, test.y)
		if total != test.n {
			t.Errorf("Can't you do simple math? (%d+%d) was incorrect, got: %d, want: %d.", test.x, test.y, test, test.n)
		}
	}
}
