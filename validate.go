package main

import "fmt"

func ValidateInput(s string) (rune, error) {
	for _, ch := range s {
		if ch < 32 || ch > 126 {
			return ch, fmt.Errorf("Invalid Character %c found!!", ch)
		}
	}
	return 0, nil
}
