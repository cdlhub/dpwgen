package internal

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode/utf8"
)

// GeneratePassword generates a password from a word list.
//
// n is the number of words the password is made of.
func GeneratePassword(wordList io.Reader, n uint) (string, error) {
	scanner := bufio.NewScanner(wordList)
	if !scanner.Scan() {
		return "", fmt.Errorf("empty word list")
	}
	text := scanner.Text()
	fields := strings.Fields(text)
	if len(fields) != 2 {
		return "", fmt.Errorf("word list bad format: must be 'nnnn password': is: %s", text)
	}

	id := fields[0]
	if _, err := strconv.ParseUint(id, 0, 64); err != nil {
		return "", fmt.Errorf("word list bad format: password id is not unsigned int: %s", id)
	}

	d := utf8.RuneCountInString(id)
	return getPassword(wordList, d)
}

func getPassword(wordList io.Reader, d int) (string, error) {

	return "", nil
}
