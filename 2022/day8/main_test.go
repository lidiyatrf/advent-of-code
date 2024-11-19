package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"advent-of-code/file/reader"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected int
	}{
		{
			name: "test_1",
			input: strings.NewReader(`30373
25512
65332
33549
35390`),
			expected: 21,
		},
		{
			name: "test_2",
			input: strings.NewReader(
				`303
625
653`),
			expected: 9,
		},
		{
			name: "test_3",
			input: strings.NewReader(
				`9999999999999999
1234555566667890
1246890999998873
9999999999999999`),
			expected: 36 + 16,
		},
		{
			name: "test_4",
			input: strings.NewReader(
				`92345555666
92463941239
92468909999
99999999999`),
			expected: 26 + 7,
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
