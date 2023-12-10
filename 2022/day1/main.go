package main

import (
	"advent-of-code/file"
	"fmt"
	"strconv"
)

func main() {
	data, err := file.ParseToStrings("2022/day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(arr []string) int {
	if len(arr) == 1 {
		return 0
	}

	var currentElfCarries int
	var elfWithMaxCaloriesCarries int

	for _, next := range arr {
		if next == "" {
			if currentElfCarries > elfWithMaxCaloriesCarries {
				elfWithMaxCaloriesCarries = currentElfCarries
			}
			currentElfCarries = 0
			continue
		}

		num, err := strconv.Atoi(next)
		if err != nil {
			panic("unsupported")
		}
		currentElfCarries += num
	}

	if currentElfCarries > elfWithMaxCaloriesCarries {
		elfWithMaxCaloriesCarries = currentElfCarries
	}

	return elfWithMaxCaloriesCarries
}

func puzzle2(arr []string) int {
	if len(arr) == 1 {
		return 0
	}

	var currentElfCarries int
	top3 := make([]int, 3)

	for _, next := range arr {
		if next == "" {
			replaced, ok := saveToTop3(top3, currentElfCarries)
			for ok {
				replaced, ok = saveToTop3(top3, replaced)
			}
			currentElfCarries = 0
			continue
		}

		num, err := strconv.Atoi(next)
		if err != nil {
			panic("unsupported")
		}
		currentElfCarries += num
	}

	replaced, ok := saveToTop3(top3, currentElfCarries)
	for ok {
		replaced, ok = saveToTop3(top3, replaced)
	}
	currentElfCarries = 0

	return top3[0] + top3[1] + top3[2]
}

func saveToTop3(top3 []int, value int) (int, bool) {
	replaced := -1
	if value > top3[0] {
		replaced = top3[0]
		top3[0] = value
	} else if value > top3[1] {
		replaced = top3[1]
		top3[1] = value
	} else if value > top3[2] {
		replaced = top3[2]
		top3[2] = value
	}
	return replaced, replaced != -1
}
