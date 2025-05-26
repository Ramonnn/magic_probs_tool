package utils

import (
	"bufio"
	"os"
	"strings"
)

// ReadCardList reads a file line-by-line, trims spaces, and returns a slice of non-empty card names.
func ReadCardList(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cards []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			cards = append(cards, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return cards, nil
}
