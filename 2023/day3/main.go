package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	data, err := file.ParseToStrings("2023/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

var (
	numbersRegexp = regexp.MustCompile("\\d+")
	symbolsRegexp = regexp.MustCompile("(!|@|#|%|&|:|;|,|/|=|-|_|\\^|[*]|[+]|[|]|[$])")
)

func puzzle1(arr []string) int {
	var res int

	for i, line := range arr {
		var prevLine string
		var nextLine string
		if i != 0 {
			prevLine = arr[i-1]
		}
		if i != len(arr)-1 {
			nextLine = arr[i+1]
		}

		lineNumbersIndexes := numbersRegexp.FindAllStringIndex(line, -1)
		for _, number := range lineNumbersIndexes {
			border := getLeftSymbol(number[0], line) + getRightSymbol(number[1], line) +
				extractSubstring(number[0], number[1], prevLine) + extractSubstring(number[0], number[1], nextLine)
			symbols := symbolsRegexp.FindAllStringIndex(border, -1)
			if len(symbols) != 0 {
				n, _ := strconv.Atoi(line[number[0]:number[1]])
				res += n
			}
		}
	}

	return res
}

func getLeftSymbol(start int, line string) string {
	if start == 0 {
		return ""
	}
	return line[start-1 : start]
}

func getRightSymbol(end int, line string) string {
	if end == len(line) {
		return ""
	}
	return line[end : end+1]
}

func extractSubstring(start, end int, line string) string {
	if line == "" {
		return ""
	}
	return getLeftSymbol(start, line) + line[start:end] + getRightSymbol(end, line)
}

var gearRegexp = regexp.MustCompile("[*]")

func puzzle2(lines []string) int {
	var res int

	for i, line := range lines {
		gearIndexes := gearRegexp.FindAllStringIndex(line, -1)
		if len(gearIndexes) == 0 {
			continue
		}

		var prevLine string
		var nextLine string
		var prevLineNumbers [][]int
		var nextLineNumbers [][]int
		if i != 0 {
			prevLine = lines[i-1]
			prevLineNumbers = numbersRegexp.FindAllStringIndex(prevLine, -1)
		}
		if i != len(lines)-1 {
			nextLine = lines[i+1]
			nextLineNumbers = numbersRegexp.FindAllStringIndex(nextLine, -1)
		}
		currentLineNumbers := numbersRegexp.FindAllStringIndex(lines[i], -1)

		for _, gear := range gearIndexes {
			nums := getNums(gear, prevLineNumbers, prevLine)
			if len(nums) > 2 {
				continue
			}
			nums = append(nums, getNums(gear, currentLineNumbers, line)...)
			if len(nums) > 2 {
				continue
			}
			nums = append(nums, getNums(gear, nextLineNumbers, nextLine)...)
			if len(nums) == 2 {
				res += nums[0] * nums[1]
			}
		}
	}

	return res
}

func getNums(gear []int, lineNumbers [][]int, line string) []int {
	var gearNumbers []int

	for _, numberLocation := range lineNumbers {
		if gear[0] >= numberLocation[0]-1 && gear[0] <= numberLocation[1] {
			number, _ := strconv.Atoi(line[numberLocation[0]:numberLocation[1]])
			gearNumbers = append(gearNumbers, number)

			continue
		}
	}

	return gearNumbers
}
