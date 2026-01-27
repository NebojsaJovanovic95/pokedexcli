package main

import (
	"strings"
	"testing"
)

// cleanInput splits the input string into words, trimming spaces.
func cleanInput(input string) []string {
	// Use strings.Fields which automatically splits by whitespace and removes extra spaces
	return strings.Fields(input)
}

// Test function
func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "foo    bar baz",
			expected: []string{"foo", "bar", "baz"},
		},
		{
			input:    "   singleword   ",
			expected: []string{"singleword"},
		},
		{
			input:    "    ",
			expected: []string{},
		},
		{
			input:    "multiple   spaces   between  words",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d, got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("For input '%s', expected word '%s' at position %d, got '%s'", c.input, c.expected[i], i, actual[i])
			}
		}
	}
}
