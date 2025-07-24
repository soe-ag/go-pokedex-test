package main

import "testing"

func TestCleanInput (t *testing.T){
cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	// add more cases here
		{
		input:    "  soe  aung aye  ",
		expected: []string{"soe", "aung", "aye"},
	},
	{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
}


for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) returned %d elements, expected %d",
				c.input, len(actual), len(c.expected))
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%q)[%d] = %q; expected %q",
					c.input, i, actual[i], c.expected[i])
			}
		}
	}
}