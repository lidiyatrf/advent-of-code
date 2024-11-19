package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"advent-of-code/file/reader"
)

func TestPuzzle1(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected int
	}{
		{
			name: "test_1",
			input: strings.NewReader(
				`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`,
			),
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			lines, err := reader.ToStrings(tc.input)
			require.NoError(t, err)

			result := puzzle1(lines)
			require.Equal(t, tc.expected, result)
		})
	}
}
