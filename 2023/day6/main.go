package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2023/day6/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

var numbersRegexp = regexp.MustCompile("\\d+")

func puzzle1(lines []string) int {
	res := 1

	times := numbersRegexp.FindAllString(lines[0], -1)
	distances := numbersRegexp.FindAllString(lines[1], -1)

	for i, time := range times {
		var waysToWin int
		distance, _ := strconv.Atoi(distances[i])

		timeMs, _ := strconv.Atoi(time)
		for holdMs := 0; holdMs < timeMs; holdMs++ {
			speed := holdMs
			mm := (timeMs - holdMs) * speed
			if mm > distance {
				waysToWin++
			}
		}

		res *= waysToWin
	}

	return res
}

func puzzle2(lines []string) int {
	var waysToWin int

	time, _ := strconv.Atoi(strings.Join(numbersRegexp.FindAllString(lines[0], -1), ""))
	distance, _ := strconv.Atoi(strings.Join(numbersRegexp.FindAllString(lines[1], -1), ""))

	for holdMs := 0; holdMs < time; holdMs++ {
		speed := holdMs
		mm := (time - holdMs) * speed
		if mm > distance {
			waysToWin++
		}
	}

	return waysToWin
}
