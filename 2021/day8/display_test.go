package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDisplay(t *testing.T) {
	input :=
		[]string{
			"acedgfb", // 8
			"cdfbe",   // 5
			"gcdfa",   // 2
			"fbcad",   // 3
			"dab",     // 7
			"cefabd",  // 9
			"cdfgeb",  // 6
			"eafb",    // 4
			"cagedb",  // 0
			"ab",      // 1,
		}
	expectedCurrLetterToCorrect := map[rune]rune{
		'd': 'a', // 100: 97
		'e': 'b', // +
		'a': 'c', // 97: 99
		'f': 'd', // 102: 100
		'g': 'e', // 103: 101
		'b': 'f', // 98: 102
		'c': 'g', // ????
	}
	expectedDigitToCurrentSignal := map[int]string{
		0: "cagedb",  // 0 +
		1: "ab",      // 1 +,
		2: "gcdfa",   // 2
		3: "fbcad",   // 3
		4: "eafb",    // 4 +
		5: "cdfbe",   // 5 +
		6: "cdfgeb",  // 6 +
		7: "dab",     // 7 +
		8: "acedgfb", // 8 +
		9: "cefabd",  // 9 +
	}

	d := newDisplay(input)
	assert.Equal(t, expectedCurrLetterToCorrect, d.currLetterToCorrect)
	assert.Equal(t, expectedDigitToCurrentSignal, d.digitToCurrentSignal)
}
