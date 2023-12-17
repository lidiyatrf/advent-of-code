package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_0",
			input: []string{
				`.|.`,
				`|..`,
				`.\|`,
			},
			expected: 7,
		},
		{
			name: "test_1",
			input: []string{
				`.\.`,
				`|..`,
				`.\.`,
			},
			expected: 5,
		},
		{
			name: "test_2",
			input: []string{
				`.\.`,
				`|.-`,
				`.\|`,
			},
			expected: 8,
		},
		{
			name: "test_3",
			input: []string{
				`.\..`,
				`|.-.`,
				`.\|.`,
				`\../`,
			},
			expected: 15,
		},
		{
			name: "test_4",
			input: []string{
				`.|...\....`,
				`.....|-...`,
				`.........\`,
				`..../.\\..`,
				`.-.-/..|..`,
				`.|....-|.\`,
				`..//.|....`,
			},
			expected: 38,
		},
		{
			name: "test_5",
			input: []string{
				`.|...\....`,
				`|.-.\.....`,
				`.....|-...`,
				`........|.`,
				`..........`,
				`.........\`,
				`..../.\\..`,
				`.-.-/..|..`,
				`.|....-|.\`,
				`..//.|....`,
			},
			expected: 46,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			require.Equal(t, tc.expected, result)
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
				`.|...\....`,
				`|.-.\.....`,
				`.....|-...`,
				`........|.`,
				`..........`,
				`.........\`,
				`..../.\\..`,
				`.-.-/..|..`,
				`.|....-|.\`,
				`..//.|....`,
			},
			expected: 51,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
