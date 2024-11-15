package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := file.ToStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

var (
	gameNumberRegexp = regexp.MustCompile("Game ([^ ]+):")
	gameRegexp       = regexp.MustCompile("[\\d+ (red|green|blue)]+")

	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
)

func puzzle1(arr []string) int {
	var res int

	for _, game := range arr {
		isGamePossible := false

		subsets := strings.Split(strings.Split(game, ":")[1], ";")
		for _, subset := range subsets {
			numberAndColors := gameRegexp.FindAllStringSubmatch(subset, -1)

			for _, numberAndColor := range numberAndColors {
				numberAndColorSplitted := strings.Split(strings.Trim(numberAndColor[0], " "), " ")
				number, _ := strconv.Atoi(numberAndColorSplitted[0])
				color := numberAndColorSplitted[1]

				switch color {
				case "green":
					isGamePossible = number <= greenLimit
				case "red":
					isGamePossible = number <= redLimit
				case "blue":
					isGamePossible = number <= blueLimit
				default:
					panic("unsupported color [" + color + "]")
				}

				if !isGamePossible {
					break
				}
			}

			if !isGamePossible {
				break
			}
		}

		if !isGamePossible {
			continue
		}

		gameNumber := gameNumberRegexp.FindStringSubmatch(game)
		gameNumberInt, _ := strconv.Atoi(gameNumber[1])

		res += gameNumberInt
	}

	return res
}

func puzzle2(arr []string) int {
	var res int

	for _, game := range arr {
		subsets := strings.Split(strings.Split(game, ":")[1], ";")

		var minimumGreen int
		var minimumRed int
		var minimumBlue int

		for _, subset := range subsets {
			numberAndColors := gameRegexp.FindAllStringSubmatch(subset, -1)

			for _, numberAndColor := range numberAndColors {
				numberAndColorSplitted := strings.Split(strings.Trim(numberAndColor[0], " "), " ")
				number, _ := strconv.Atoi(numberAndColorSplitted[0])
				color := numberAndColorSplitted[1]

				switch color {
				case "green":
					if number > minimumGreen {
						minimumGreen = number
					}
				case "red":
					if number > minimumRed {
						minimumRed = number
					}
				case "blue":
					if number > minimumBlue {
						minimumBlue = number
					}
				default:
					panic("unsupported color [" + color + "]")
				}
			}
		}

		powerOfTheMinimumSet := minimumGreen * minimumRed * minimumBlue
		res += powerOfTheMinimumSet
	}

	return res
}
