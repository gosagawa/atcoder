package main

import (
	"fmt"
	"testing"
)

func TestSortForFive(t *testing.T) {

	tests := []struct {
		waight   []int
		expected string
	}{
		{
			expected: "ABCDE",
			waight:   []int{1, 2, 3, 4, 5},
		},
		{
			expected: "BACDE",
			waight:   []int{2, 1, 3, 4, 5},
		},
		{
			expected: "ABECD",
			waight:   []int{1, 2, 4, 5, 3},
		},
		{
			expected: "EDCBA",
			waight:   []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			c := &TesterComparison{
				waight: tt.waight,
			}
			fmt.Println("-------")
			fmt.Println(tt.expected)
			result := SortForFive(c, []rune("ABCDE"))
			if string(result) != tt.expected {
				t.Fatalf("return want to be %+v but returned %+v", tt.expected, string(result))
			}
		})
	}
}

type TesterComparison struct {
	waight []int
}

func (c *TesterComparison) Gt(a, b rune) bool {
	return c.waight[runeToInt(a)] > c.waight[runeToInt(b)]
}
func runeToInt(r rune) int {
	s := string(r)
	switch s {
	case "A":
		return 0
	case "B":
		return 1
	case "C":
		return 2
	case "D":
		return 3
	case "E":
		return 4
	}
	return 0
}
