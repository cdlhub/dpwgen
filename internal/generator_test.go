package internal

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	var tests = []struct {
		input string
		n     uint
	}{
		{create5DiceWordListToString(), 1},
		{create5DiceWordListToString(), 10},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("GeneratePassword(%q, %d)", test.input, test.n)
		r1 := strings.NewReader(test.input)
		pw1, err := GeneratePassword(r1, test.n)
		if err != nil {
			t.Errorf("%s: should pass: failed: %v", descr, err)
		}
		if len(strings.Fields(pw1)) != int(test.n) {
			t.Errorf("%s: wrong number of word: want: %d: got: %d", descr, test.n, len(strings.Fields(pw1)))
		}
		r2 := strings.NewReader(test.input)
		pw2, _ := GeneratePassword(r2, test.n)
		if pw1 == pw2 {
			t.Errorf("%s: bad seed (or bad luck): same password generated twice: %q", descr, pw1)
		}
	}
}

func TestLoadWordList(t *testing.T) {
	var tests = []struct {
		input string
		n     int
		ok    bool
	}{
		{"1 un\ndeux 2\n3 trois\n4 quatre\n5 cinq\n6 six", 1, false},
		{"1 un\n2\n3 trois\n4 quatre\n5 cinq\n6 six", 1, false},
		{"1 un\n2 deux et demi\n3 trois\n4 quatre\n5 cinq\n6 six", 1, false},
		{createOneDieWordListToString(), 1, true},
		{create5DiceWordListToString(), 5, true},
		{createWordListWithTooFewWordsToString(), 3, false},
		{createWordListWithWrongIDToString(), 3, false},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("loadWordList(%q)", test.input)
		r := strings.NewReader(test.input)
		if _, n, err := loadWordList(r); (err == nil) != test.ok {
			t.Errorf("%s: want: (%d, %t): got: (%d, %t): %v", descr, test.n, test.ok, n, err == nil, err)
		} else if (err == nil) && test.ok {
			if test.n != n {
				t.Errorf("%s: wrong number of dice: want %d: got: %d", descr, test.n, n)
			}
		}
	}
}

func TestCheckWordListLength(t *testing.T) {
	var tests = []struct {
		inputList map[string]string
		inputN    int
		ok        bool
	}{
		{nil, 5, false},
		{map[string]string{}, 5, false},
		{map[string]string{"1": "un", "2": "deux", "3": "trois", "4": "quatre", "5": "cinq"}, 1, false},
		{createOneDieWordList(), 1, true},
		{map[string]string{"1": "un", "2": "deux", "3": "trois", "4": "quatre", "5": "cinq", "6": "six", "7": "sept"}, 1, false},
		{map[string]string{"1": "un", "2": "deux", "3": "trois", "4": "quatre", "5": "cinq", "6": "six"}, 6, false},
		{map[string]string{"1": "un"}, 1, false},
		{create5DiceWordList(), 5, true},
	}

	for _, test := range tests {
		if err := checkWordListLength(test.inputList, test.inputN); (err == nil) != test.ok {
			t.Errorf("checkWordListLength(%v, %d): want: %t: got: %t: %v", test.inputList, test.inputN, test.ok, err == nil, err)
		}
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

func TestGetRandomSide(t *testing.T) {
	descr := fmt.Sprintf("getRandomSide()")

	rand.Seed(1)
	m := make(map[int]int)
	for i := uint16(0); i < math.MaxUint16; i++ {
		r := getRandomSide()
		m[r]++
	}
	for i := 1; i <= DIESIDES; i++ {
		if _, ok := m[i]; !ok {
			t.Errorf("%s: not all %d sides generated", descr, DIESIDES)
		}
	}

	var ok bool
	for k := range m {
		ok = false
		for i := 1; i <= DIESIDES; i++ {
			if k == i {
				ok = true
				break
			}
		}
		if !ok {
			t.Errorf("%s: wrong die side: got: %d", descr, k)
		}
	}
}

func TestGetRandomWord(t *testing.T) {
	var tests = []struct {
		input map[string]string
		n     uint
	}{
		{createOneDieIDList(), 1},
		{create5DiceWordList(), 5},
	}

	rand.Seed(1)
	for _, test := range tests {
		descr := fmt.Sprintf("getRandomWord(%d)", test.n)
		got := make(map[string]string)
		for i := 0; i < 100*len(test.input); i++ {
			w, err := getRandomWord(test.input, test.n)
			if err != nil {
				t.Errorf("%s: should succeed: failed: %v", descr, err)
			}
			got[w] = w
		}
		if !reflect.DeepEqual(got, test.input) {
			for k := range test.input {
				if _, ok := got[k]; !ok {
					t.Errorf("%s: missing key: %s", descr, k)
				}
			}
			t.Errorf("%s: generated words mismatch: want length: %v: got length: %v", descr, len(test.input), len(got))
		}
	}
}

func create5DiceWordList() map[string]string {
	var wl = make(map[string]string)
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
	return wl
}

func createOneDieIDList() map[string]string {
	var ids = make(map[string]string)
	for i := 1; i <= DIESIDES; i++ {
		ids[strconv.Itoa(i)] = strconv.Itoa(i)
	}
	return ids
}

func createOneDieWordList() map[string]string {
	return map[string]string{"1": "un", "2": "deux", "3": "trois", "4": "quatre", "5": "cinq", "6": "six"}
}

func createOneDieWordListToString() string {
	return mapToString(createOneDieWordList())
}

func create5DiceWordListToString() string {
	return mapToString(create5DiceWordList())
}

func createWordListWithTooFewWordsToString() string {
	wl := create5DiceWordList()
	delete(wl, "11111")
	return mapToString(wl)
}

func createWordListWithWrongIDToString() string {
	wl := create5DiceWordList()
	delete(wl, "11111")
	wl["11117"] = "11117"
	return mapToString(wl)
}

func mapToString(m map[string]string) string {
	var s string
	for k, v := range m {
		s += k + " " + v + "\n"
	}
	return s
}
