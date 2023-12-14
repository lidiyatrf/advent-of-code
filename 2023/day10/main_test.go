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
			name: "test_1",
			input: []string{
				".....",
				".S-7.",
				".|.|.",
				".L-J.",
				".....",
			},
			expected: 4,
		},
		{
			name: "test_1",
			input: []string{
				"..F7.",
				".FJ|.",
				"SJ.L7",
				"|F--J",
				"LJ...",
			},
			expected: 8,
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
				".....",
				".F-7.",
				".S.|.",
				".L-J.",
				".....",
			},
			expected: 1,
		},
		{
			name: "test_1",
			input: []string{
				"...........",
				".F--S----7.",
				".|F-----7|.",
				".||.....||.",
				".||.....||.",
				".|L-7.F-J|.",
				".|..|.|..|.",
				".L--J.L--J.",
				"...........",
			},
			expected: 4,
		},
		{
			name: "test_2",
			input: []string{
				"..........",
				".F--S---7.",
				".|F----7|.",
				".||....||.",
				".||....||.",
				".|L-7F-J|.",
				".|..||..|.",
				".L--JL--J.",
				"..........",
			},
			expected: 4,
		},
		{
			name: "test_3",
			input: []string{
				".F----7F7F7F7F-7....",
				".|F--7||||||||FJ....",
				".||.FJ||||||||L7....",
				"FJL7L7LJLJ||LJ.L-7..",
				"L--J.L7...LJF7F-7L7.",
				"....F-J..F7FJ|L7L7L7",
				"....L7.F7||L7S.L7L7|",
				".....|FJLJ|FJ|F7|.LJ",
				"....FJL-7.||.||||...",
				"....L---J.LJ.LJLJ...",
			},
			expected: 8,
		},
		{
			name: "test_4",
			input: []string{
				"FF7F7F7F7F7F7F7F---7",
				"L|LJS|||||||||||F--J",
				"FL-7LJLJ||||||LJL-77",
				"F--JF--7||LJLJ7F7FJ-",
				"L---JF-JLJ.||-FJLJJ7",
				"|F|F-JF---7F7-L7L|7|",
				"|FFJF7L7F-JF7|JL---7",
				"7-L-JL7||F7|L7F-7F7|",
				"L.L7LFJ|||||FJL7||LJ",
				"L7JLJL-JLJLJL--JLJ.L",
			},
			expected: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
