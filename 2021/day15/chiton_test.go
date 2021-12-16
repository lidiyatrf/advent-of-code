package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLowestRiskPath(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 0",
			input: []string{
				"1",
				"2",
			},
			result: 2,
		},
		{
			name: "test 1",
			input: []string{
				"116",
				"138",
			},
			result: 12,
		},
		{
			name: "test 2",
			input: []string{
				"1163",
				"1381",
				"2136",
			},
			result: 13,
		},
		{
			name: "test 3",
			input: []string{
				"112",
				"131",
				"683",
				"316",
			},
			result: 13,
		},
		{
			name: "test 4",
			input: []string{
				"1163751742",
				"1381373672",
				"2136511328",
				"3694931569",
				"7463417111",
				"1319128137",
				"1359912421",
				"3125421639",
				"1293138521",
				"2311944581",
			},
			result: 40,
		},
		{
			name: "test 5",
			input: []string{
				"1199999999",
				"9119999999",
				"9911999999",
				"9991199999",
				"9999119999",
				"9999911999",
				"9999991199",
				"9999999119",
				"9999999911",
			},
			result: 17,
		},
		{
			name: "test 6",
			input: []string{
				"1999999999",
				"1999999999",
				"1999999999",
				"1999999999",
				"1999999999",
				"1999999999",
				"1999999999",
				"1111111111",
			},
			result: 16,
		},
		{
			name: "test 7",
			input: []string{
				"199",
				"119",
				"919",
				"119",
				"199",
				"111",
			},
			result: 9,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getLowestRiskPath(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}
