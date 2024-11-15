package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		distictSymbols int
		expected       int
	}{
		{
			name:           "puzzle1_1",
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			distictSymbols: 4,
			expected:       7,
		},
		{
			name:           "puzzle1_2",
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			distictSymbols: 4,
			expected:       5,
		},
		{
			name:           "puzzle1_3",
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			distictSymbols: 4,
			expected:       6,
		},
		{
			name:           "puzzle1_4",
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			distictSymbols: 4,
			expected:       10,
		},
		{
			name:           "puzzle1_5",
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			distictSymbols: 4,
			expected:       11,
		},
		{
			name:           "puzzle2_1",
			input:          "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			distictSymbols: 14,
			expected:       19,
		},
		{
			name:           "puzzle2_2",
			input:          "bvwbjplbgvbhsrlpgdmjqwftvncz",
			distictSymbols: 14,
			expected:       23,
		},
		{
			name:           "puzzle2_3",
			input:          "nppdvjthqldpwncqszvftbrmjlhg",
			distictSymbols: 14,
			expected:       23,
		},
		{
			name:           "puzzle2_4",
			input:          "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			distictSymbols: 14,
			expected:       29,
		},
		{
			name:           "puzzle2_5",
			input:          "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			distictSymbols: 14,
			expected:       26,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle(tc.input, tc.distictSymbols)
			require.Equal(t, tc.expected, result)
		})
	}
}
