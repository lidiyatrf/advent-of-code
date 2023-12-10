package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
)

func main() {
	data, err := file.ParseToStrings("2023/day3/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
}

var (
	numbersRegexp = regexp.MustCompile("\\d+")
	symbolsRegexp = regexp.MustCompile("(!|@|#|%|&|:|;|,|/|=|-|_|\\^|[*]|[+]|[|]|[$])")
)

func puzzle1(arr []string) int {
	var res int

	linePrev := arr[0]
	linePrevNumbersIndexes := numbersRegexp.FindAllStringIndex(linePrev, -1)
	linePrevNumbersFlags := make([]bool, len(linePrevNumbersIndexes))

	for _, line := range arr[1:] {
		lineSymbolsIndexes := symbolsRegexp.FindAllStringIndex(line, -1)
		fmt.Println(symbolsRegexp)
		res += len(line)
	}

	return res
}
