package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathsCount(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 1",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			result: 10,
		},
		{
			name: "test 2",
			input: []string{
				"start-A",
				"start-B",
				"A-end",
				"B-end",
			},
			result: 2,
		},
		{
			name: "test 3",
			input: []string{
				"dc-end",
				"HN-start",
				"start-kj",
				"dc-start",
				"dc-HN",
				"LN-dc",
				"HN-end",
				"kj-sa",
				"kj-HN",
				"kj-dc",
			},
			result: 19,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getPathsCount(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}
