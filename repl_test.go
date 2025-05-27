package main

import (
	"fmt"
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
		expected: nil,
	}, {
		input:    "HellOWorld",
		expected: []string{"helloworld"},
	}}

	for _, c := range cases {
		actual := cleanImput(c.input)

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("%s does not equal %s. Test Failed", word, expectedWord)
			}

		}
		fmt.Println("Test passed")
	}
}
