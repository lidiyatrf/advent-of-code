package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetPathsCountPart1(t *testing.T) {
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
			result := getPathsCountPart1(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestGetPathsCountPart2(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		result int
	}{
		{
			name: "test 0",
			input: []string{
				"start-a",
				"start-b",
				"a-b",
				"a-end",
				"b-end",
			},
			result: 6,
		},
		{
			name: "test 1",
			input: []string{
				"start-A",
				"A-b",
				"A-end",
			},
			result: 3,
		},
		{
			name: "test 2",
			input: []string{
				"start-b",
				"b-d",
				"b-end",
			},
			result: 2,
		},
		{
			name: "test 3",
			input: []string{
				"start-b",
				"start-A",
				"b-d",
				"A-b",
				"A-end",
				"b-end",
			},
			result: 13,
		},
		{
			name: "example 1",
			input: []string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			result: 36,
		},
		{
			name: "example 2",
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
			result: 103,
		},
		{
			name: "example 3",
			input: []string{
				"fs-end",
				"he-DX",
				"fs-he",
				"start-DX",
				"pj-DX",
				"end-zg",
				"zg-sl",
				"zg-pj",
				"pj-he",
				"RW-he",
				"fs-DX",
				"pj-RW",
				"zg-RW",
				"start-pj",
				"he-WI",
				"zg-he",
				"pj-fs",
				"start-RW",
			},
			result: 3509,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getPathsCountPart2(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}
