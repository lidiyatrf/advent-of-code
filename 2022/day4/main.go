package main

import (
	"advent-of-code/2021/file"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2022/day4/input.txt")
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
		ranges := strings.Split(line, ",")

		start1, end1 := splitRange(ranges[0])
		start2, end2 := splitRange(ranges[1])

		if firstRangeContainsSecond(start1, end1, start2, end2) ||
			firstRangeContainsSecond(start2, end2, start1, end1) {
			res++
			continue
		}
	}

	return res
}

func firstRangeContainsSecond(start1 int, end1 int, start2 int, end2 int) bool {
	return start1 <= start2 && end1 >= end2
}

func splitRange(value string) (start int, end int) {
	split := strings.Split(value, "-")
	start, err := strconv.Atoi(split[0])
	if err != nil {
		panic("unsupported")
	}
	end, err = strconv.Atoi(split[1])
	if err != nil {
		panic("unsupported")
	}
	return
}

func puzzle2(arr []string) int {
	var res int

	for _, line := range arr {
		ranges := strings.Split(line, ",")

		start1, end1 := splitRange(ranges[0])
		start2, end2 := splitRange(ranges[1])

		if firstRangeOverlapsSecond(start1, end1, start2, end2) ||
			firstRangeOverlapsSecond(start2, end2, start1, end1) {
			res++
			continue
		}
	}

	return res
}

func firstRangeOverlapsSecond(start1 int, end1 int, start2 int, end2 int) bool {
	return end1 >= start2 && end2 >= start1
}
