package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Go   Programming   Language  ",
			expected: []string{"go", "programming", "language"},
		},
		{
			input:    "ONE\tTWO\nTHREE\rFOUR",
			expected: []string{"one", "two", "three", "four"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "SingleWord",
			expected: []string{"singleword"},
		},
		{
			input:    "  MiXeD   CaSe   StRiNg  ",
			expected: []string{"mixed", "case", "string"},
		},
		{
			input:    "\t\n  TABS   AND\r\nNEWLINES  \t\n",
			expected: []string{"tabs", "and", "newlines"},
		},
		{
			input:    "multiple    spaces     between",
			expected: []string{"multiple", "spaces", "between"},
		},
		{
			input:    "  A  B  C  D  E  ",
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			input:    "Numbers123 Mixed456 With789 Text",
			expected: []string{"numbers123", "mixed456", "with789", "text"},
		},
		{
			input:    "UPPER lower MiXeD",
			expected: []string{"upper", "lower", "mixed"},
		},
		{
			input:    "   \t\n\r   ",
			expected: []string{},
		},
		{
			input:    "Special!@# Characters$%^ Here&*()",
			expected: []string{"special!@#", "characters$%^", "here&*()"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Length of actual slice produced by 'cleanInput': %v does not match length of expected slice: %v.", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word produced by cleanInput: %v does not match word in expected: %v", word, expectedWord)
			}
		}
	}
}
