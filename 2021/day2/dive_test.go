package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMove(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expectedResult := 150

	actualResult, err := move(input)
	require.NoError(t, err)
	require.Equal(t, expectedResult, actualResult)
}

func TestGetShift(t *testing.T) {
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

func TestMoveWithAim(t *testing.T) {
	input := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	expectedResult := 900

	actualResult, err := moveWithAim(input)
	require.NoError(t, err)
	require.Equal(t, expectedResult, actualResult)
}
