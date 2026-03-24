package main

import (
	"strings"
	"testing"
)

// helper to count lines
func lineCount(s string) int {
	if s == "" {
		return 0
	}

	lines := strings.Split(s, "\n")

	// remove last empty line caused by final \n
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	return len(lines)
}

func TestPrintASCII(t *testing.T) {

	// create fake banner lines (same size as real banner)
	mockBanner := make([]string, 9*95)

	for i := range mockBanner {
		mockBanner[i] = "X"
	}

	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Single letter",
			input:    "A",
			expected: 8,
		},
		{
			name:     "Word",
			input:    "HELLO",
			expected: 8,
		},
		{
			name:     "With newline",
			input:    "A\\nB",
			expected: 16,
		},
		{
			name:     "Empty input",
			input:    "",
			expected: 0,
		},
		{
			name:     "Only newline",
			input:    "\\n",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := printASCII(tt.input, mockBanner)

			lines := lineCount(result)

			if lines != tt.expected {
				t.Errorf("expected %d lines but got %d", tt.expected, lines)
			}
		})
	}
}
