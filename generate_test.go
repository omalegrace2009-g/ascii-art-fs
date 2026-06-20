package main

import (
	"strings"
	"testing"
)

func mockBanner() map[rune][]string {
	return map[rune][]string{
		'A': {
			"A1", "A2", "A3", "A4",
			"A5", "A6", "A7", "A8",
		},
		'B': {
			"B1", "B2", "B3", "B4",
			"B5", "B6", "B7", "B8",
		},
	}
}
func TestGenerateArt_EmptyInput(t *testing.T) {
	got := GenerateArt("", mockBanner())

	if got != "" {
		t.Errorf("expected empty string, got %q", got)
	}
}
func TestGenerateArt_OnlyNewline(t *testing.T) {
	got := GenerateArt("\\n", mockBanner())

	if got != "\n" {
		t.Errorf("expected one newline, got %q", got)
	}
}
func TestGenerateArt_MultipleNewlines(t *testing.T) {
	got := GenerateArt("\\n\\n\\n", mockBanner())

	if got != "\n\n\n" {
		t.Errorf("expected 3 newlines, got %q", got)
	}
}
func TestGenerateArt_SingleCharacter(t *testing.T) {
	got := GenerateArt("A", mockBanner())

	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

	if len(lines) != 8 {
		t.Fatalf("expected 8 lines, got %d", len(lines))
	}

	if lines[0] != "A1" {
		t.Errorf("expected A1, got %q", lines[0])
	}
}
func TestGenerateArt_MultipleCharacters(t *testing.T) {
	got := GenerateArt("AB", mockBanner())

	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

	if lines[0] != "A1B1" {
		t.Errorf("expected A1B1, got %q", lines[0])
	}
}
func TestGenerateArt_MultipleLines(t *testing.T) {
	got := GenerateArt("A\\nB", mockBanner())

	lines := strings.Split(strings.TrimRight(got, "\n"), "\n")

	if len(lines) != 16 {
		t.Fatalf("expected 16 lines, got %d", len(lines))
	}
}
func TestGenerateArt_EmptyLineBetweenText(t *testing.T) {
	got := GenerateArt("A\\n\\nB", mockBanner())

	if !strings.Contains(got, "\n\n") {
		t.Errorf("expected blank line between rendered text")
	}
}
func TestGenerateArt_LeadingNewline(t *testing.T) {
	got := GenerateArt("\\nA", mockBanner())

	if got[0] != '\n' {
		t.Errorf("expected output to start with newline")
	}
}
func TestGenerateArt_TrailingNewline(t *testing.T) {
	got := GenerateArt("A\\n", mockBanner())

	if !strings.HasSuffix(got, "\n") {
		t.Errorf("expected trailing newline")
	}
}
func TestGenerateArt_UnknownCharacter(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("GenerateArt panicked on unknown character")
		}
	}()

	GenerateArt("@", mockBanner())
}
