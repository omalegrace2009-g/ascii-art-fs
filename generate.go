package main

import "strings"

func GenerateArt(s string, banner map[rune][]string) string {
	if s == "" {
		return ""
	}

	var result strings.Builder
	line := strings.Split(s, "\\n")
	if strings.Join(line, "") == "" {
		return strings.Repeat("\n", len(line)-1)
	}
	
	for _, g := range line {
		if g == "" {
			result.WriteString("\n")
		} else {
			rend := RenderLine(g, banner)
			for _, r := range rend {
				result.WriteString(r)
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
