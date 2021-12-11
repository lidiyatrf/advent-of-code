package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimulate(t *testing.T) {
	input :=
		[]string{
			"11111",
			"19991",
			"19191",
			"19991",
			"11111",
		}
	expectedResult1 := [][]int{
		{3, 4, 5, 4, 3},
		{4, 0, 0, 0, 4},
		{5, 0, 0, 0, 5},
		{4, 0, 0, 0, 4},
		{3, 4, 5, 4, 3},
	}
	expectedResult2 := [][]int{
		{4, 5, 6, 5, 4},
		{5, 1, 1, 1, 5},
		{6, 1, 1, 1, 6},
		{5, 1, 1, 1, 5},
		{4, 5, 6, 5, 4},
	}
	expectedFlashed := 9

	actualFlashed := 0
	b := newBoard(input)
	b.simulate()
	actualFlashed += b.getJustFlashed()
	require.Equal(t, expectedResult1, b.levels)
	b.simulate()
	actualFlashed += b.getJustFlashed()
	require.Equal(t, expectedResult2, b.levels)
	require.Equal(t, expectedFlashed, actualFlashed)
}

func TestFlash(t *testing.T) {
	input :=
		[]string{
			"111",
			"199",
			"191",
		}
	expectedResult := [][]int{
		{2, 2, 2},
		{2, 10, 10},
		{2, 10, 2},
	}

	b := newBoard(input)
	b.flash(1, 1)
	require.Equal(t, expectedResult, b.levels)
}
