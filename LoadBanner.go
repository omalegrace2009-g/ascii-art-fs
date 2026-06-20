package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	ban, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	line := strings.Split(string(ban), "\n")
	banner := make(map[rune][]string)
	ch := rune(32)
	for i := 0; i < len(line); i += 9 {
		if i+8 < len(line) {
			banner[ch] = line[i+1 : i+9]
			ch++
		}
	}
	if len(banner) == 0 {
		return nil, errors.New("Error Found!!")
	}
	return banner, err
}
