package main

import (
	"fmt"
	"testing"
)

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{4, 6},
		{-1, 1},
		{99999, 100001},
	}

	for _, test := range tests {
		if output := Calc(test.input); output != test.expected {
			t.Error(fmt.Sprintf("Test failed: %v inputted, %v expected, %v received", test.input, test.expected, output))
		}
	}
}
