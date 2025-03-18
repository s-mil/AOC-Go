// Skeleton from https://github.com/alexchao26/advent-of-code-go/blob/main/scripts/skeleton/tmpls/main_test.go
package main

import (
	"testing"
)

var example = ``

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "present1",
			input: "2x3x4",
			want:  58,
		},
		{
			name:  "present2",
			input: "1x1x10",
			want:  43,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Present1",
			input: "2x3x4",
			want:  34,
		},
		{
			name:  "Present2",
			input: "1x1x10",
			want:  14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
