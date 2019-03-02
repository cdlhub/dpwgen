package internal

import (
	"strconv"
	"testing"
)

func TestCheckWordListLength(t *testing.T) {
	var tests = []struct {
		inputList map[string]string
		inputN    int
		ok        bool
	}{
		{nil, 5, false},
		{map[string]string{}, 5, false},
		{map[string]string{"1": "un", "2": "deux", "3": "trois", "4": "quatre", "5": "cinq", "6": "six"}, 1, true},
		{map[string]string{"1": "un", "2": "deux", "3": "trois", "4": "quatre", "5": "cinq", "6": "six"}, 6, false},
		{map[string]string{"1": "un"}, 1, false},
	}

	for _, test := range tests {
		if err := checkWordListLength(test.inputList, test.inputN); (err == nil) != test.ok {
			t.Errorf("checkWordListLength(%v, %d): want: %t: got: %t: %v", test.inputList, test.inputN, test.ok, err == nil, err)
		}
	}

	var wl = make(map[string]string)
	var n int
	for d1 := 1; d1 <= 6; d1++ {
		for d2 := 1; d2 <= 6; d2++ {
			for d3 := 1; d3 <= 6; d3++ {
				for d4 := 1; d4 <= 6; d4++ {
					for d5 := 1; d5 <= 6; d5++ {
						id := strconv.Itoa(d1) + strconv.Itoa(d2) + strconv.Itoa(d3) + strconv.Itoa(d4) + strconv.Itoa(d5)
						wl[id] = id
					}
				}
			}
		}
	}
	if err := checkWordListLength(wl, 5); err != nil {
		t.Errorf("checkWordListLength(%q, %d): want: %t: got: %t: %v", wl, n, true, err == nil, err)
	}

}

func TestCheckID(t *testing.T) {
	var tests = []struct {
		inputID string
		inputN  int
		ok      bool
	}{
		{"", 5, false},
		{"4", 1, true},
		{"8", 1, false},
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
