package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
				"#.#.### 1,1,3",
				".#...#....###. 1,1,3",
				".#.###.#.###### 1,3,1,6",
				"####.#...#... 4,1,1",
				"#....######..#####. 1,6,5",
				".###.##....# 3,2,1",
			},
			expected: 21,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := puzzle1(tc.input)
			assert.Equal(t, tc.expected, result)
		})
	}
}
