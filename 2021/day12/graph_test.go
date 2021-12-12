package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewGraph(t *testing.T) {
	input :=
		[]string{
			"start-A",
			"start-b",
			"A-c",
			"A-b",
			"b-d",
			"A-end",
			"b-end",
		}
	expectedResult := 6

	g := newGraph(input)
	actualResult := len(g.vertexes)
	require.Equal(t, expectedResult, actualResult)
}
