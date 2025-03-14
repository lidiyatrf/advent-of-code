package main

import (
	"advent-of-code/file/reader"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected int
	}{
		{
			name: "test_1",
			input: strings.NewReader(
				`111
222`),
			expected: 6,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lines, err := reader.ToStrings(tc.input)
			require.NoError(t, err)

			result := puzzle1(lines)
			require.Equal(t, tc.expected, result)
		})
	}
}
