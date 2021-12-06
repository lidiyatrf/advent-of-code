package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSchool(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result map[int]int
	}{
		{
			name:  "test 1",
			input: "3,4,3,1,2",
			result: func() map[int]int {
				m := make(map[int]int)
				m[0] = 0
				m[1] = 1
				m[2] = 1
				m[3] = 2
				m[4] = 1
				m[5] = 0
				m[6] = 0
				m[7] = 0
				m[8] = 0
				return m
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			school, err := newSchool(tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.result, school.fishesPerDay)
		})
	}
}

func TestGrow(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result map[int]int
	}{
		{
			name:  "test 1",
			input: "3,4,3,1,2",
			result: func() map[int]int {
				m := make(map[int]int)
				m[0] = 1
				m[1] = 1
				m[2] = 2
				m[3] = 1
				m[4] = 0
				m[5] = 0
				m[6] = 0
				m[7] = 0
				m[8] = 0
				return m
			}(),
		},
		{
			name:  "test 1",
			input: "0,1,2,2,3",
			result: func() map[int]int {
				m := make(map[int]int)
				m[0] = 1
				m[1] = 2
				m[2] = 1
				m[3] = 0
				m[4] = 0
				m[5] = 0
				m[6] = 1
				m[7] = 0
				m[8] = 1
				return m
			}(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			school, err := newSchool(tc.input)
			require.NoError(t, err)
			school.grow()
			require.Equal(t, tc.result, school.fishesPerDay)
		})
	}
}

func TestTotal(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result int
	}{
		{
			name:   "test 1",
			input:  "3,4,3,1,2",
			result: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			school, err := newSchool(tc.input)
			require.NoError(t, err)
			total := school.getTotal()
			require.Equal(t, tc.result, total)
		})
	}
}
