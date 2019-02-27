package internal

import "os"

// GeneratePassword generates a password from word list file fileName.
//
// d is the number of dice to roll to chose a word from the list.
//
// n is the number of words the password is made of.
func GeneratePassword(fileName string, d, n uint) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	return "", nil
}
