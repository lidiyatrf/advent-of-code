package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
				"",
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			expected: 405,
		},
		{
			name: "test_1",
			input: []string{
				"##....###.#..",
				"########...#.",
				"##.##.##.#.##",
				"........###.#",
				"........#.#.#",
				"#######.###..",
				".#.##.#.#....",
				"..####..##..#",
				"##.##.##.##.#",
				"##.##.##.##.#",
				"..####..##..#",
				".#.##.#.#....",
				"#######.###..",
				"........#.#.#",
				"........###.#",
				"",
				"#.###..#..###",
				".#...##.####.",
				".#...##.####.",
				"#.###..#..###",
				".#######.##.#",
				".#..##.#.#..#",
				"..#..#.##.#..",
				"##..##..###.#",
				"######.##..#.",
				"######.##....",
				"##..##..###.#",
				"..#..#.##.#..",
				".#..##.#.#..#",
				"",
				"#......",
				"#......",
				"....#..",
				"#.#.###",
				"#####..",
				"###.##.",
				"...#.##",
				"#.#.###",
				"####.#.",
				"..##..#",
				"..##..#",
				"####.#.",
				"#.#.###",
				"...#.##",
				"###.##.",
				"###.#..",
				"#.#.###",
			},
			expected: 1200,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFindVerticalMirror(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			expected: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, ok := findVerticalMirror(tc.input)
			require.Equal(t, true, ok)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestFindHorizontalMirror(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			expected: 4,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, ok := findHorizontalMirror(tc.input)
			require.Equal(t, true, ok)
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
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
				"",
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			expected: 400,
		},
		{
			name: "vert1",
			input: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"....##..#",
				"..##..##.",
				"#.#.##.#.",
			},
			expected: 300,
		},
		{
			name: "vert2",
			input: []string{
				"##....###.#..",
				"########...#.",
				"##.##.##.#.##",
				"........###.#",
				"........#.#.#",
				"#######.###..",
				".#.##.#.#....",
				"..####..##..#",
				"##.##.##.##.#",
				"##.##.##.##.#",
				"..####..##..#",
				".#.##.#.#....",
				"#######.###..",
				"........#.#.#",
				"........###.#",
			},
			expected: 1400,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
