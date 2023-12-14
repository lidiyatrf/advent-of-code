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
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: 136,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestCalcBoard(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: 104,
		},
		{
			name: "test_2",
			input: []string{
				".....#....",
				"....#...O#",
				".....##...",
				"...#......",
				".....OOO#.",
				".O#...O#.#",
				"....O#...O",
				"......OOOO",
				"#....###.O",
				"#.OOO#..OO",
			},
			expected: 64,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := calcBoard(stringsToRune(tc.input))
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
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: 64,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle2(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestOneCycle(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: []string{
				".....#....",
				"....#...O#",
				"...OO##...",
				".OO#......",
				".....OOO#.",
				".O#...O#.#",
				"....O#....",
				"......OOOO",
				"#...O###..",
				"#..OO#....",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runeToStrings(oneCycle(stringsToRune(tc.input)))
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestTwoCycles(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: []string{
				".....#....",
				"....#...O#",
				".....##...",
				"..O#......",
				".....OOO#.",
				".O#...O#.#",
				"....O#...O",
				".......OOO",
				"#..OO###..",
				"#.OOO#...O",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runeToStrings(oneCycle(oneCycle(stringsToRune(tc.input))))
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestThreeCycles(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: []string{
				".....#....",
				"....#...O#",
				".....##...",
				"..O#......",
				".....OOO#.",
				".O#...O#.#",
				"....O#...O",
				".......OOO",
				"#...O###.O",
				"#.OOO#...O",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runeToStrings(oneCycle(oneCycle(oneCycle(stringsToRune(tc.input)))))
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestRollSouth(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: []string{
				".....#....",
				"....#....#",
				"...O.##...",
				"...#......",
				"O.O....O#O",
				"O.#..O.#.#",
				"O....#....",
				"OO....OO..",
				"#OO..###..",
				"#OO.O#...O",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runeToStrings(rollSouth(stringsToRune(tc.input)))
			require.Equal(t, tc.expected, result)
		})
	}
}
func TestRollNorth(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: []string{
				"OOOO.#.O..",
				"OO..#....#",
				"OO..O##..O",
				"O..#.OO...",
				"........#.",
				"..#....#.#",
				"..O..#.O.O",
				"..O.......",
				"#....###..",
				"#....#....",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runeToStrings(rollNorth(stringsToRune(tc.input)))
			require.Equal(t, tc.expected, result)
		})
	}
}
func TestRollWest(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: []string{
				"O....#....",
				"OOO.#....#",
				".....##...",
				"OO.#OO....",
				"OO......#.",
				"O.#O...#.#",
				"O....#OO..",
				"O.........",
				"#....###..",
				"#OO..#....",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runeToStrings(rollWest(stringsToRune(tc.input)))
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestRollEast(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name: "test_1",
			input: []string{
				"O....#....",
				"O.OO#....#",
				".....##...",
				"OO.#O....O",
				".O.....O#.",
				"O.#..O.#.#",
				"..O..#O..O",
				".......O..",
				"#....###..",
				"#OO..#....",
			},
			expected: []string{
				"....O#....",
				".OOO#....#",
				".....##...",
				".OO#....OO",
				"......OO#.",
				".O#...O#.#",
				"....O#..OO",
				".........O",
				"#....###..",
				"#..OO#....",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := runeToStrings(rollEast(stringsToRune(tc.input)))
			require.Equal(t, tc.expected, result)
		})
	}
}
