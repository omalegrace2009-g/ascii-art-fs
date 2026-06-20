package main

func RenderLine(s string, banner map[rune][]string) []string {
	result := make([]string, 8)
	for i := 0; i <= 7; i++ {
		for _, ch := range s {
			result[i] += banner[ch][i]
		}
	}
	return result
}
