package main

import (
	"testing"
)

func Test_NotQuiteLisp(t *testing.T) {
	tests := []struct {
		input  string
		part   int
		output int
	}{
		{"(())", 1, 0},
		{"()()", 1, 0},
		{"(((", 1, 3},
		{"(()(()(", 1, 3},
		{"))(((((", 1, 3},
		{"())", 1, -1},
		{"))(", 1, -1},
		{")))", 1, -3},
		{")())())", 1, -3},
		{")",2,1},
		{"()())",2,5},
	}

	for _, test := range tests {
		if got := NotQuiteLisp(test.input, test.part); got != test.output {
			t.Errorf("NotQuiteLisp(%s) = %d; want %d", test.input, got, test.output)
		}
	}
}
