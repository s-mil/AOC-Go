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
			name:  "one",
			input: ">",
			want:  2,
		},
		{
			name:  "two",
			input: "^>v<",
			want:  4,
		},
		{
			name:  "three",
			input: "^v^v^v^v^v",
			want:  2,
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
			name:  "one",
			input: "^v",
			want:  3,
		},
		{
			name:  "two",
			input: "^>v<",
			want:  3,
		},
		{
			name:  "three",
			input: "^v^v^v^v^v",
			want:  11,
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
