package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	data, err := file.ToStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data)) //53897
}

var digitRegexp = regexp.MustCompile("\\d")

func puzzle1(arr []string) int {
	var res int

	for _, line := range arr {
		numbers := digitRegexp.FindAllString(line, -1)
		if len(numbers) == 0 {
			panic("no numbers in line")
		}

		firstDigit, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(fmt.Sprintf("error to parse digit [%s]", numbers[0]))
		}
		if len(numbers) == 1 {
			number := firstDigit*10 + firstDigit
			res += number
			continue
		}

		secondDigit, err := strconv.Atoi(numbers[len(numbers)-1])
		if err != nil {
			panic(fmt.Sprintf("error to parse digit [%s]", numbers[len(numbers)-1]))
		}
		number := firstDigit*10 + secondDigit
		res += number
	}

	return res
}

var (
	numbersRegexp         = regexp.MustCompile("([\\d]|one|two|three|four|five|six|seven|eight|nine)")
	numbersReversedRegexp = regexp.MustCompile("([\\d]|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)")
)

func puzzle2(arr []string) int {
	var res int

	for _, line := range arr {
		firstNumber, _ := strDigitToInt(numbersRegexp.FindString(line))
		secondNumber, _ := strDigitToInt(Reverse(numbersReversedRegexp.FindString(Reverse(line))))

		number := firstNumber*10 + secondNumber
		res += number
	}

	return res
}

func strDigitToInt(str string) (int, error) {
	switch str {
	case "one":
		return 1, nil
	case "two":
		return 2, nil
	case "three":
		return 3, nil
	case "four":
		return 4, nil
	case "five":
		return 5, nil
	case "six":
		return 6, nil
	case "seven":
		return 7, nil
	case "eight":
		return 8, nil
	case "nine":
		return 9, nil
	default:
		return strconv.Atoi(str)
	}
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
