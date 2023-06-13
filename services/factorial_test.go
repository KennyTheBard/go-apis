package services

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	var tests = []struct {
		n   int
		exp int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%v!", tt.n)
		t.Run(testName, func(t *testing.T) {
			result := Factorial(tt.n)
			if result != tt.exp {
				t.Errorf("Expected %v, but received %v", tt.exp, result)
			}
		})
	}
}
