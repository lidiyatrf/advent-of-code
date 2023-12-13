package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 35,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestPuzzle2(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"seeds: 79 14 55 13",
				"",
				"seed-to-soil map:",
				"50 98 2",
				"52 50 48",
				"",
				"soil-to-fertilizer map:",
				"0 15 37",
				"37 52 2",
				"39 0 15",
				"",
				"fertilizer-to-water map:",
				"49 53 8",
				"0 11 42",
				"42 0 7",
				"57 7 4",
				"",
				"water-to-light map:",
				"88 18 7",
				"18 25 70",
				"",
				"light-to-temperature map:",
				"45 77 23",
				"81 45 19",
				"68 64 13",
				"",
				"temperature-to-humidity map:",
				"0 69 1",
				"1 0 69",
				"",
				"humidity-to-location map:",
				"60 56 37",
				"56 93 4",
			},
			expected: 46,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
