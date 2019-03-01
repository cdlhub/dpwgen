package internal

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	// DIESIDES is the number of side of simulated die
	DIESIDES = 6
	// SIDESET is the set of possible side numbers for a word ID
	SIDESET = "[123456]"
)

// GeneratePassword generates a password from a word list.
//
// n is the number of words the passphrase is made of.
func GeneratePassword(wordList io.Reader, n uint) (string, error) {
	wl, d, err := loadWordList(wordList)
	if err != nil {
		return "", fmt.Errorf("generate password: %v", err)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	var pw []string
	for i := uint(0); i < n; i++ {
		w, err := getRandomWord(wl, uint(d))
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
		if !checkID(fields[0], nDice) {
			return nil, -1, fmt.Errorf("word list: line %d: ID bad format: should be composed of %d numbers from 1 to %d", i, nDice, DIESIDES)
		}

		wl[fields[0]] = fields[1]
	}

	if err := checkWordListLength(wl, nDice); err != nil {
		return wl, nDice, fmt.Errorf("word list: %v", err)
	}

	return wl, nDice, nil
}

// n is number of dice
func checkWordListLength(wl map[string]string, n int) error {
	nWordsWanted := int(math.Pow(float64(DIESIDES), float64(n)))
	if len(wl) != nWordsWanted {
		return fmt.Errorf("wrong number of words: want: %d: got: %d", nWordsWanted, len(wl))
	}
	return nil
}

// n is number of dice
func checkID(id string, n int) bool {
	r := regexp.MustCompile(fmt.Sprintf("^%s+$", SIDESET))
	return r.MatchString(id) && utf8.RuneCountInString(id) == n
}

func getRandomWord(wl map[string]string, d uint) (string, error) {
	var id string
	for i := uint(0); i < d; i++ {
		id += strconv.Itoa(rand.Intn(DIESIDES-1) + 1)
	}

	w, ok := wl[id]
	if !ok {
		return w, fmt.Errorf("get word: id %q not found", id)
	}

	return w, nil
}
