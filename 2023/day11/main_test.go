package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int64
	}{
		{
			name: "test_1",
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			expected: 374,
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
		name                              string
		input                             []string
		replacementForEmptyRowsAndColumns int64
		expected                          int64
	}{
		{
			name: "test_10",
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			replacementForEmptyRowsAndColumns: 10,
			expected:                          1030,
		},
		{
			name: "test_100",
			input: []string{
				"...#......",
				".......#..",
				"#.........",
				"..........",
				"......#...",
				".#........",
				".........#",
				"..........",
				".......#..",
				"#...#.....",
			},
			replacementForEmptyRowsAndColumns: 100,
			expected:                          8410,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle(tc.input, tc.replacementForEmptyRowsAndColumns)
			assert.Equal(t, tc.expected, result)
		})
	}
}
