package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name: "test_0",
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3",
				"",
			},
			expected: "NDP",
		},
		{
			name: "test_1",
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3",
				"",
				"move 1 from 2 to 1",
			},
			expected: "DCP",
		},
		{
			name: "test_4",
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3",
				"",
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			expected: "CMZ",
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
		expected string
	}{
		{
			name: "test_1",
			input: []string{
				"    [D]    ",
				"[N] [C]    ",
				"[Z] [M] [P]",
				" 1   2   3",
				"",
				"move 1 from 2 to 1",
				"move 3 from 1 to 3",
				"move 2 from 2 to 1",
				"move 1 from 1 to 2",
			},
			expected: "MCD",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
