package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 4361,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
				"......#...",
				"617*......",
				".....+.58.",
				"..592.....",
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 467835,
		},
		{
			name: "test_2",
			input: []string{
				"467..114..",
				"...*......",
				"..35..633.",
			},
			expected: 16345,
		},
		{
			name: "test_3",
			input: []string{
				"......755.",
				"...$.*....",
				".664.598..",
			},
			expected: 451490,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
