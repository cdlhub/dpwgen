package internal

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
	"unicode/utf8"
)

// GeneratePassword generates a password from a word list.
//
// n is the number of words the password is made of.
func GeneratePassword(wordList io.Reader, n uint) (string, error) {
	wl, d, err := loadWordList(wordList)
	if err != nil {
		return "", fmt.Errorf("generate password: %v", err)
	}

	var pw []string
	for i := uint(0); i < n; i++ {
		w, err := getRandomWord(wl, d)
		if err != nil {
			return w, fmt.Errorf("generate password: %v", err)
		}
		pw = append(pw, w)
	}

	return strings.Join(pw, " "), nil
}

func loadWordList(wordList io.Reader) (map[string]string, int, error) {
	wl := make(map[string]string)
	nDice := -1

	scanner := bufio.NewScanner(wordList)
	for i := uint(1); scanner.Scan(); i++ {
		text := scanner.Text()
		fields := strings.Fields(text)
		if len(fields) != 2 {
			return nil, -1, fmt.Errorf("word list: line %d: bad format: must be 'nnnn password': is: %s", i, text)
		}

		d := utf8.RuneCountInString(fields[0])
		if nDice == -1 {
			nDice = d
		}
		if err := checkID(fields[0], nDice); err != nil {
			return nil, -1, fmt.Errorf("word list: line %d: %v", i, err)
		}

		wl[fields[0]] = fields[1]
	}

	return wl, nDice, nil
}

func checkID(id string, n int) error {
	r := regexp.MustCompile("^[123456]+$")
	if !r.MatchString(id) {
		return fmt.Errorf("id bad format: should be composed of 1, 2, 3, 4, 5, or 6 only: is: %s", id)
	}

	if utf8.RuneCountInString(id) != n {
		return fmt.Errorf("id bad format: should be of length %d: length is: %d", n, len(id))
	}

	return nil
}

func getRandomWord(wl map[string]string, d int) (string, error) {

	return "", nil
}
