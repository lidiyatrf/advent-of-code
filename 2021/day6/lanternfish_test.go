package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountLanternfish(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
		days   int
	}{
		{
			name:   "test 1",
			input:  "3,4,3,1,2",
			days:   1,
			result: 5,
		},
		{
			name:   "test 1",
			input:  "3,4,3,1,2",
			days:   2,
			result: 6,
		},
		{
			name:   "test 1",
			input:  "3,4,3,1,2",
			days:   3,
			result: 7,
		},
		{
			name:   "test 1",
			input:  "3,4,3,1,2",
			days:   18,
			result: 26,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result, err := countLanternfish(tc.input, tc.days)
			require.NoError(t, err)
			require.Equal(t, tc.result, result)
		})
	}
}
