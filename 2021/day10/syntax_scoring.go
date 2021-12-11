package main

import (
	"fmt"

	"advent-of-code/2021/file"
)

var openClosePairs = map[rune]rune{
	'{': '}',
	'[': ']',
	'<': '>',
	'(': ')',
}

func main() {
	data, err := file.ParseToStrings("2021/day10/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	score := calcScoreForSyntaxErrors(data)
	if err != nil {
		fmt.Println(score)
		return
	}
	fmt.Println("score:", score)
}

func calcScoreForSyntaxErrors(data []string) int {
	sum := 0
	for _, nextRow := range data {
		char := getFirstIllegalCharInRow(nextRow)
		switch char {
		case ')':
			sum += 3
		case ']':
			sum += 57
		case '}':
			sum += 1197
		case '>':
			sum += 25137
		}
	}
	return sum
}

func getFirstIllegalCharInRow(row string) rune {
	var stack []rune
	for _, nextBracket := range row {
		if _, isOpening := openClosePairs[nextBracket]; isOpening {
			stack = append(stack, nextBracket)
			continue
		}
		if nextBracket == openClosePairs[stack[len(stack)-1]] {
			stack = stack[0 : len(stack)-1]
			continue
		}
		return nextBracket
	}
	// ignoring all the others
	return '-'
}
