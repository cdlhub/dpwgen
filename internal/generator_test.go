package internal

import (
	"testing"
)

func TestCheckID(t *testing.T) {
	var tests = []struct {
		inputID string
		inputN  int
		ok      bool
	}{
		{"", -1, false},
		{"", 0, false},
		{"", 5, false},
		{"4", 1, true},
		{"8", 1, false},
		{"123", -1, false},
		{"123", 0, false},
		{"123", 2, false},
		{"123", 3, true},
		{"123", 4, false},
		{"0123", 4, false},
		{"1237", 4, false},
		{"1283", 4, false},
		{"1246", 4, true},
		{"3333", 4, true},
		{"123456", 6, true},
		{"632113", 6, true},
	}

	for _, test := range tests {
		if ok := checkID(test.inputID, test.inputN); ok != test.ok {
			t.Errorf("checkID(%v, %d): want: %t: got: %t", test.inputID, test.inputN, test.ok, ok)
		}
	}
}
