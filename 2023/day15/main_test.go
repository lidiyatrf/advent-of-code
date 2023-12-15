package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "test_0",
			input:    "HASH",
			expected: 52,
		},
		{
			name:     "test_1",
			input:    "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			expected: 1320,
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
		input    string
		expected int
	}{
		{
			name:     "test_1",
			input:    "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7",
			expected: 145,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
