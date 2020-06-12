package main

import (
	"fmt"
	"testing"
)


func TestParseArgs(t *testing.T) {
	var tests = []struct {
		input    []string
		expected string
		err      error
	}{
		{input: []string{}, err: fmt.Errorf("usage: lower [string]")},
		{input: []string{"foo"}, expected: "foo"},
		{input: []string{"foo", "bar"}, err: fmt.Errorf("expected exactly one argument, found 2")},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("parseArgs(%s)", test.input), func(t *testing.T) {
			actual, err := parseArgs(test.input)
			if err != nil {
				if test.err != nil {
					if err.Error() != test.err.Error() {
						t.Errorf("Expected error: %s, but go %s", test.err.Error(), err.Error())
					}
				} else {
					t.Errorf("Unexpected error: %s", err.Error())
				}
			} else {
				if test.err != nil {
					t.Errorf("Expected error: %s to occur", test.err.Error())
				}
				if actual != test.expected {
					t.Errorf("Expected %s but got %s", test.expected, actual)
				}
			}
		})
	}
}
