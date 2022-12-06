package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, test := range tests {
		if ans := checkMarker(test.input, 4); ans != test.answer {
			t.Errorf("Got %d, wanted %d\n", ans, test.answer)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input  string
		answer int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, test := range tests {
		if ans := checkMarker(test.input, 14); ans != test.answer {
			t.Errorf("Got %d, wanted %d\n", ans, test.answer)
		}
	}
}

func BenchmarkPart1(b *testing.B) {
	tests := []struct {
		input  string
		answer int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
	}

	for _, test := range tests {
		b.Run(fmt.Sprintf("input_4_%s", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				checkMarker(test.input, 4)
			}
		})
	}
}

func BenchmarkPart2(b *testing.B) {
	tests := []struct {
		input  string
		answer int
	}{
		{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, test := range tests {
		b.Run(fmt.Sprintf("input_14_%s", test.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				checkMarker(test.input, 14)
			}
		})
	}
}
