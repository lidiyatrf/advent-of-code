package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSumLowPoints(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 1",
			input: []string{
				"2199943210",
				"3987894921",
				"9856789892",
				"8767896789",
				"9899965678",
			},
			result: 15,
		},
		{
			name: "test 2",
			input: []string{
				"9876789198765442",
				"9868997346999875",
				"8754398757997654",
				"6543239898996543",
				"5432123999987954",
				"7521012999899876",
			},
			result: 26,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := sumRiskedLevel(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}
