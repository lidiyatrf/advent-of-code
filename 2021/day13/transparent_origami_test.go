package main

import (
	"testing"
)

func TestGetPointsAfterFirstFold(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 0",
			input: []string{
				"6,10",
				"0,14",
				"9,10",
				"0,3",
				"10,4",
				"4,11",
				"6,0",
				"6,12",
				"4,1",
				"0,13",
				"10,12",
				"3,4",
				"3,0",
				"8,4",
				"1,10",
				"2,14",
				"8,10",
				"9,0",
				"",
				"fold along y=7",
				"fold along x=5",
			},
			result: 17,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			//result := getPointsAfterFirstFold(tc.input)
			//require.Equal(t, tc.result, result)

			printPointsAfterAllFolds(tc.input)
		})
	}
}
