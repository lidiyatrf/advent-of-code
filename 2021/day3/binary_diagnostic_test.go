package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLifeSupportRating(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 1",
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			result: 230,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getLifeSupportRating(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestGetCo2ScrubberRating(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result string
	}{
		{
			name: "test 1",
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			result: "01010",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getCo2ScrubberRating(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestGetOxygenGeneratorRating(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result string
	}{
		{
			name: "test 1",
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			result: "10111",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getOxygenGeneratorRating(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestFindMostCommonBit(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		position int
		result   rune
	}{
		{
			name: "more 1",
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
			},
			position: 0,
			result:   '1',
		},
		{
			name: "more 0",
			input: []string{
				"11110",
				"10110",
				"10111",
				"10101",
			},
			position: 1,
			result:   '0',
		},
		{
			name: "equal amount",
			input: []string{
				"00100",
				"11110",
			},
			position: 0,
			result:   '1',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := findMostCommonBit(tc.input, tc.position)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestFindLeastCommonBit(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		position int
		result   rune
	}{
		{
			name: "more 1",
			input: []string{
				"00100",
				"11110",
				"10110",
				"10111",
			},
			position: 0,
			result:   '0',
		},
		{
			name: "more 0",
			input: []string{
				"11110",
				"10110",
				"10111",
				"10101",
			},
			position: 1,
			result:   '1',
		},
		{
			name: "equal amount",
			input: []string{
				"00100",
				"11110",
			},
			position: 0,
			result:   '0',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := findLeastCommonBit(tc.input, tc.position)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestGetPowerConsumption(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name:   "test 1",
			input:  []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			result: 198,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := getPowerConsumption(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestMustConvertBinToDec(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
	}{
		{
			name:   "21",
			input:  "0010101",
			result: 21,
		},
		{
			name:   "119",
			input:  "1110111",
			result: 119,
		},
		{
			name:   "0",
			input:  "0000000",
			result: 0,
		},
		{
			name:   "1",
			input:  "0000001",
			result: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mustConvertBinToDec(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}
