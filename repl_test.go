package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input  string
		output []string
	}{
		{
			input:  "Hello world",
			output: []string{"hello", "world"},
		},
		{
			input:  "Get to the choppa",
			output: []string{"get", "to", "the", "choppa"},
		},
		{
			input:  "",
			output: []string{},
		},
	}

	for _, c := range cases {
		result := cleanInput(c.input)
		if len(result) != len(c.output) {
			t.Errorf("Output of cleanInput doesn't produce correct number of outputs")
		}

		// make sure each word matches
		for i := range result {
			if result[i] != c.output[i] {
				t.Errorf("Word %v doesn't match expected word %v. Index: %v", result[i], c.output[i], i)
			}
		}
	}

}
