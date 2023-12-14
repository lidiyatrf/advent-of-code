package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
)

func main() {
	data, err := file.ParseToStrings("2023/day4/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

var partsRegexp = regexp.MustCompile(`: (.*?) \| (.*?)$`)
var numbersRegexp = regexp.MustCompile(`\d+`)

func puzzle1(lines []string) int {
	var res int

	for _, line := range lines {
		parts := partsRegexp.FindStringSubmatch(line)

		winningNumbers := make(map[string]struct{})

		for _, n := range numbersRegexp.FindAllString(parts[1], -1) {
			winningNumbers[n] = struct{}{}
		}

		var elfsNumbersCounter int
		for _, n := range numbersRegexp.FindAllString(parts[2], -1) {
			if _, ok := winningNumbers[n]; ok {
				if elfsNumbersCounter == 0 {
					elfsNumbersCounter = 1
					continue
				}
				elfsNumbersCounter *= 2
			}
		}
		res += elfsNumbersCounter
	}

	return res
}

func puzzle2(lines []string) int {
	var res int
	amountOfCards := make([]int, len(lines))

	for cardNumber, line := range lines {
		parts := partsRegexp.FindStringSubmatch(line)

		winningNumbers := make(map[string]struct{})
		for _, n := range numbersRegexp.FindAllString(parts[1], -1) {
			winningNumbers[n] = struct{}{}
		}

		var winningCardsCounter int
		for _, n := range numbersRegexp.FindAllString(parts[2], -1) {
			if _, ok := winningNumbers[n]; ok {
				winningCardsCounter++
			}
		}

		amountOfCards[cardNumber] += 1
		for j := 0; j < winningCardsCounter; j++ {
			amountOfCards[cardNumber+j+1] += amountOfCards[cardNumber]
		}

		res += amountOfCards[cardNumber]
	}
	return res
}
