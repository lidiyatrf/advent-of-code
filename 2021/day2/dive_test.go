package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMove_Positive(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "test 1",
			input:    []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			expected: 150,
		},
	}

	for _, tc := range tests {
		t.Run(strings.Join(tc.input, ", "), func(t *testing.T) {
			result, err := move(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMove_Err(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "unsupported action",
			input: []string{"down 5", "left 8"},
		},
		{
			name:  "cannot parse amount",
			input: []string{"forward", "down 5asd", "forward 2"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := move(tc.input)
			require.Error(t, err)
		})
	}
}

func TestGetShift_Positive(t *testing.T) {
	tests := []struct {
		input     string
		expectedX int
		expectedY int
	}{
		{
			input:     "forward 5",
			expectedX: 5,
			expectedY: 0,
		},
		{
			input:     "down 5",
			expectedX: 0,
			expectedY: 5,
		},
		{
			input:     "up 3",
			expectedX: 0,
			expectedY: -3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			x, y, err := getShift(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expectedX, x)
			require.Equal(t, tc.expectedY, y)
		})
	}
}

func TestGetShift_Err(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			input: "left 5",
		},
		{
			input: "up aa",
		},
		{
			input: "up 3aa",
		},
	}

	for _, tc := range tests {
		t.Run(tc.input, func(t *testing.T) {
			_, _, err := getShift(tc.input)
			require.Error(t, err)
		})
	}
}
