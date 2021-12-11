package main

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetFirstIllegalCharInRow(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		result rune
	}{
		{
			name:   "test 1",
			input:  "{([(<{}[<>[]}>{[]{[(<()>",
			result: '}',
		},
		{
			name:   "test 2",
			input:  "[[<[([]))<([[{}[[()]]]",
			result: ')',
		},
		{
			name:   "test 3",
			input:  "[{[{({}]{}}([{[{{{}}([]",
			result: ']',
		},
		{
			name:   "test 4",
			input:  "[<(<(<(<{}))><([]([]()",
			result: ')',
		},
		{
			name:   "test 5",
			input:  "<{([([[(<>()){}]>(<<{{",
			result: '>',
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := getFirstIllegalCharInRow(tc.input)
			require.Equal(t, tc.result, result)
		})
	}
}

func TestCalcScoreForSyntaxErrors(t *testing.T) {
	input :=
		[]string{
			"[({(<(())[]>[[{[]{<()<>>",
			"[(()[<>])]({[<{<<[]>>(",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"(((({<>}<{<{<>}{[]{[]{}",
			"[[<[([]))<([[{}[[()]]]",
			"[{[{({}]{}}([{[{{{}}([]",
			"{<[[]]>}<{[{[{[]{()[[[]",
			"[<(<(<(<{}))><([]([]()",
			"<{([([[(<>()){}]>(<<{{",
			"<{([{{}}[<[[[<>{}]]]>[]]",
		}
	expectedResult := 26397

	actualResult := calcScoreForSyntaxErrors(input)
	assert.Equal(t, expectedResult, actualResult)
}

func TestCalcScoreForIncompleteRows(t *testing.T) {
	input :=
		[]string{
			"[({(<(())[]>[[{[]{<()<>>",
			"[(()[<>])]({[<{<<[]>>(",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"(((({<>}<{<{<>}{[]{[]{}",
			"[[<[([]))<([[{}[[()]]]",
			"[{[{({}]{}}([{[{{{}}([]",
			"{<[[]]>}<{[{[{[]{()[[[]",
			"[<(<(<(<{}))><([]([]()",
			"<{([([[(<>()){}]>(<<{{",
			"<{([{{}}[<[[[<>{}]]]>[]]",
		}
	expectedResult := 288957

	actualResult := calcScoreForIncompleteRows(input)
	assert.Equal(t, expectedResult, actualResult)
}
