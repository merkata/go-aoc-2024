package day1_test

import (
	"testing"
	"github.com/merkata/go-aoc-2024/day1"
	"github.com/merkata/go-aoc-2024/common"
)

func TestSolve1(t *testing.T) {
	tests := []struct{
		filepath string
		want int
	}{
		{"day1_example.txt", 11},
		{"day1_input.txt", 1889772},
	}
	for _, tt := range tests {
		fixture, _ := common.InputToStrings(tt.filepath)
		got := day1.Solve1(fixture)
		if got != tt.want {
			t.Errorf("Got %d, want %d", got, tt.want)
		}
	} 
}

func TestSolve2(t *testing.T) {
	tests := []struct{
		filepath string
		want int
	}{
		{"day1_example.txt", 31},
		{"day1_input.txt", 23228917},
	}
	for _, tt := range tests {
		fixture, _ := common.InputToStrings(tt.filepath)
		got := day1.Solve2(fixture)
		if got != tt.want {
			t.Errorf("Got %d, want %d", got, tt.want)
		}
	} 
}
