package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetStraightOverlappedAmount(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 1",
			input: []string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			result: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getStraightOverlappedAmount(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestGetOverlappedAmount(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 1",
			input: []string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			result: 12,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getOverlappedAmount(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}
