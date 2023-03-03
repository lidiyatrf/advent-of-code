package main

import (
	"advent-of-code/2021/file"
	"fmt"
)

func main() {
	data, err := file.ParseToStrings("2022/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(arr []string) int {
	var res int

	for _, line := range arr {
		letter := getDuplicatedLetterPuzzle1(line)
		res += getPriority(letter)
	}

	return res
}

func getDuplicatedLetterPuzzle1(value string) int {
	for i := 0; i < len(value)/2; i++ {
		for j := len(value) / 2; j < len(value); j++ {
			if value[i] == value[j] {
				return int(value[i])
			}
		}
	}
	panic("unsupported")
}

func getPriority(letter int) int {
	if letter >= 'a' && letter <= 'z' {
		return letter - 'a' + 1
	}
	if letter >= 'A' && letter <= 'Z' {
		return letter - 'A' + 27
	}
	return 0
}

func puzzle2(arr []string) int {
	var res int

	for i := 0; i < len(arr); i += 3 {
		letter := getDuplicatedLetterPuzzle2(arr[i], arr[i+1], arr[i+2])
		res += getPriority(letter)
	}

	return res
}

func getDuplicatedLetterPuzzle2(line1, line2, line3 string) int {
	dupl1 := make(map[uint8]int)
	for i := 0; i < len(line1); i++ {
		dupl1[line1[i]] = 1
	}
	dupl2 := make(map[uint8]int)
	for i := 0; i < len(line2); i++ {
		dupl2[line2[i]] = 1
	}
	for i := 0; i < len(line3); i++ {
		_, firstHas := dupl1[line3[i]]
		_, SecondHas := dupl2[line3[i]]
		if firstHas && SecondHas {
			return int(line3[i])
		}
	}
	panic("unsupported")
}
