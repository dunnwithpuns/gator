package main

import (
	"testing"
	// "os"
	"fmt"
)

func TestExitCodes(t *testing.T) {
	cases := []struct {
		input []string
	}{
		{
			input: []string{"go", "run", "."},
		},
		{
			input: []string{"go", "run", ".", "login"},
		},
		{
			input: []string{"go", "run", ".", "login", "alice"},
		},
	}

	for i := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			return
		})
	}
}
