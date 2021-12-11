package main

import (
	"fmt"
	"sort"

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

	syntaxErrors := calcScoreForSyntaxErrors(data)
	fmt.Println("syntaxErrors:", syntaxErrors)

	incompleteRows := calcScoreForIncompleteRows(data)
	fmt.Println("incompleteRows:", incompleteRows)
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

func calcScoreForIncompleteRows(data []string) int {
	var scores []int

	for _, nextRow := range data {
		brackets, isIncomplete := getBracketsToComplete(nextRow)
		if !isIncomplete {
			continue
		}
		score := 0
		for _, bracket := range brackets {
			score *= 5
			switch bracket {
			case ')':
				score += 1
			case ']':
				score += 2
			case '}':
				score += 3
			case '>':
				score += 4
			}
		}
		scores = insertToSorted(scores, score)
	}
	return scores[len(scores)/2]
}

func getBracketsToComplete(row string) (brackets []rune, isIncomplete bool) {
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
		return nil, false
	}

	var result []rune
	for i := len(stack) - 1; i >= 0; i-- {
		result = append(result, openClosePairs[stack[i]])
	}
	return result, true
}

func insertToSorted(arr []int, newVal int) []int {
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= newVal })

	if i == len(arr) {
		return append(arr, newVal)
	}

	arr = append(arr[:i+1], arr[i:]...)
	arr[i] = newVal
	return arr
}
