package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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
				"R 6 (#70c710)",
				"D 5 (#0dc571)",
				"L 2 (#5713f0)",
				"D 2 (#d2c081)",
				"R 2 (#59c680)",
				"D 2 (#411b91)",
				"L 5 (#8ceee2)",
				"U 2 (#caa173)",
				"L 1 (#1b58a2)",
				"U 2 (#caa171)",
				"R 2 (#7807d2)",
				"U 3 (#a77fa3)",
				"L 2 (#015232)",
				"U 2 (#7a21e3)",
			},
			expected: 62,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int64
	}{
		{
			name: "test_1",
			input: []string{
				"R 6 (#70c710)",
				"D 5 (#0dc571)",
				"L 2 (#5713f0)",
				"D 2 (#d2c081)",
				"R 2 (#59c680)",
				"D 2 (#411b91)",
				"L 5 (#8ceee2)",
				"U 2 (#caa173)",
				"L 1 (#1b58a2)",
				"U 2 (#caa171)",
				"R 2 (#7807d2)",
				"U 3 (#a77fa3)",
				"L 2 (#015232)",
				"U 2 (#7a21e3)",
			},
			expected: 952408144115,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
