package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountFuelToAlignCrabs(t *testing.T) {
	input := "16,1,2,0,4,2,7,1,2,14"
	expectedResult := 37

	actualResult, err := countFuelToAlignCrabs(input)
	require.NoError(t, err)
	require.Equal(t, expectedResult, actualResult)
}

func TestGetMedian(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		result int
	}{
		{
			name:   "test 1",
			input:  []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			result: 2,
		},
		{
			name:   "test 2",
			input:  []int{3, 5, 1, 5, 3, 2, 1, 3, 4, 2, 5, 1, 3, 3, 2, 5, 1, 3, 1, 5, 5, 1, 1, 1, 2, 4, 1, 4, 5, 2, 1, 2, 4, 3, 1, 2, 3, 4, 3, 4, 4, 5, 1, 1, 1, 1, 5, 5, 3, 4, 4, 4, 5, 3, 4, 1, 4, 3, 3, 2, 1, 1, 3, 3, 3, 2, 1, 3, 5, 2, 3, 4, 2, 5, 4, 5, 4, 4, 2, 2, 3, 3, 3, 3, 5, 4, 2, 3, 1, 2, 1, 1, 2, 2, 5, 1, 1, 4, 1, 5, 3, 2, 1, 4, 1, 5, 1, 4, 5, 2, 1, 1, 1, 4, 5, 4, 2, 4, 5, 4, 2, 4, 4, 1, 1, 2, 2, 1, 1, 2, 3, 3, 2, 5, 2, 1, 1, 2, 1, 1, 1, 3, 2, 3, 1, 5, 4, 5, 3, 3, 2, 1, 1, 1, 3, 5, 1, 1, 4, 4, 5, 4, 3, 3, 3, 3, 2, 4, 5, 2, 1, 1, 1, 4, 2, 4, 2, 2, 5, 5, 5, 4, 1, 1, 5, 1, 5, 2, 1, 3, 3, 2, 5, 2, 1, 2, 4, 3, 3, 1, 5, 4, 1, 1, 1, 4, 2, 5, 5, 4, 4, 3, 4, 3, 1, 5, 5, 2, 5, 4, 2, 3, 4, 1, 1, 4, 4, 3, 4, 1, 3, 4, 1, 1, 4, 3, 2, 2, 5, 3, 1, 4, 4, 4, 1, 3, 4, 3, 1, 5, 3, 3, 5, 5, 4, 4, 1, 2, 4, 2, 2, 3, 1, 1, 4, 5, 3, 1, 1, 1, 1, 3, 5, 4, 1, 1, 2, 1, 1, 2, 1, 2, 3, 1, 1, 3, 2, 2, 5, 5, 1, 5, 5, 1, 4, 4, 3, 5, 4, 4},
			result: 3,
		},
		{
			name:   "test 3",
			input:  []int{3, 5, 1, 5, 3, 2, 1, 3, 4, 2},
			result: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getMedian(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestCountFuelToAlign(t *testing.T) {
	tests := []struct {
		name         string
		inputArr     []int
		inputMediana int
		result       int
	}{
		{
			name:         "test 1",
			inputArr:     []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			inputMediana: 2,
			result:       37,
		},
		{
			name:         "test 2",
			inputArr:     []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			inputMediana: 1,
			result:       41,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := countCostToAlign(tc.inputArr, tc.inputMediana)
			require.Equal(t, tc.result, result)
		})
	}
}
