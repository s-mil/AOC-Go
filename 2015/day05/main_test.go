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
			name:  "1",
			input: "ugknbfddgicrmopn",
			want:  1,
		},
		{
			name:  "2",
			input: "jchzalrnumimnmhp",
			want:  0,
		},
		{
			name:  "3",
			input: "haegwjzuvuyypxyu",
			want:  0,
		},
		{
			name:  "4",
			input: "dvszwmarrgswjxmb",
			want:  0,
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
			name:  "1",
			input: "qjhvhtzxzqqjkmpb",
			want:  1,
		},
		{
			name:  "2",
			input: "xxyxx",
			want:  1,
		},
		{
			name:  "3",
			input: "uurcxstgmygtbstg",
			want:  0,
		},
		{
			name:  "4",
			input: "ieodomkazucvgmuy",
			want:  0,
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
