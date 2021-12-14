package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetDiffMostAndLeastCommon(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		steps  int
		result int64
	}{
		{
			name: "test 0",
			input: []string{
				"NNCB",
				"",
				"CH -> B",
				"HH -> N",
				"CB -> H",
				"NH -> C",
				"HB -> C",
				"HC -> B",
				"HN -> C",
				"NN -> C",
				"BH -> H",
				"NC -> B",
				"NB -> B",
				"BN -> B",
				"BB -> N",
				"BC -> B",
				"CC -> N",
				"CN -> C",
			},
			steps:  10,
			result: 1588,
		},
		{
			name: "test 1",
			input: []string{
				"NNCB",
				"",
				"CH -> B",
				"HH -> N",
				"CB -> H",
				"NH -> C",
				"HB -> C",
				"HC -> B",
				"HN -> C",
				"NN -> C",
				"BH -> H",
				"NC -> B",
				"NB -> B",
				"BN -> B",
				"BB -> N",
				"BC -> B",
				"CC -> N",
				"CN -> C",
			},
			steps:  40,
			result: 2188189693529,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getDiffMostAndLeastCommon(tc.input, tc.steps)
			require.Equal(t, tc.result, result)
		})
	}
}
