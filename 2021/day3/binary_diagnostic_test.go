package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
