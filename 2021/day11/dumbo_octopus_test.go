package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountTotalFlashes(t *testing.T) {
	{
		tests := []struct {
			name   string
			input  []string
			steps  int
			result int
		}{
			{
				name: "test 1",
				input: []string{
					"11111",
					"19991",
					"19191",
					"19991",
					"11111",
				},
				steps:  2,
				result: 9,
			},
			{
				name: "test 2",
				input: []string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				},
				steps:  1,
				result: 0,
			},
			{
				name: "test 3",
				input: []string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				},
				steps:  2,
				result: 35,
			},
			{
				name: "test 4",
				input: []string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				},
				steps:  10,
				result: 204,
			},
			{
				name: "test 5",
				input: []string{
					"5483143223",
					"2745854711",
					"5264556173",
					"6141336146",
					"6357385478",
					"4167524645",
					"2176841721",
					"6882881134",
					"4846848554",
					"5283751526",
				},
				steps:  100,
				result: 1656,
			},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				result := countTotalFlashes(tc.input, tc.steps)
				require.Equal(t, tc.result, result)
			})
		}
	}
}

func Test(t *testing.T) {
	input := []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}
	expectedResult := 195

	actualResult := findWhenAllOctopusesFlash(input)
	require.Equal(t, expectedResult, actualResult)
}
