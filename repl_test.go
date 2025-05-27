package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{{
		input:    " hello world ",
		expected: []string{"hello", "world"},
	}, {
		input:    "HELLO WORLD",
		expected: []string{"hello", "world"},
	}, {
		input:    "",
		expected: []string{},
	}, {
		input:    "HellOWorld",
		expected: []string{"helloworld"},
	}, {
		input:    "\tHellO World\n",
		expected: []string{"hello", "world"},
	}}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected word count is %d. Actual worrd count is %d. Test failed", len(c.expected), len(actual))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%s does not equal %s. Test Failed", word, expectedWord)
				t.Fail()
			}

		}

	}
}
