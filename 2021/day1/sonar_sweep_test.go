package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountIncreasedTriple(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test_1",
			input:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expected: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := countIncreasedTriple(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestCountIncreasedSingle(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "test_1",
			input:    []int{199, 200},
			expected: 1,
		},
		{
			name:     "test_2",
			input:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expected: 7,
		},
		{
			name:     "test_3",
			input:    []int{10, 7, 8, 9, 11, 10, 20, 17, 20, 27},
			expected: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := countIncreasedSingle(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
