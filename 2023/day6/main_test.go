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
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expected: 288,
		},
		{
			name: "test_1",
			input: []string{
				"Time:      7",
				"Distance:  9",
			},
			expected: 4,
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
				"Time:      7  15   30",
				"Distance:  9  40  200",
			},
			expected: 71503,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
