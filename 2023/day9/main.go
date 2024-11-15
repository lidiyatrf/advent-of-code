package main

import (
	"advent-of-code/file"
	"fmt"
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

func puzzle1(lines []string) int {
	var res int

	for _, line := range lines {
		numbersStr := strings.Split(line, " ")

		var differences [][]int

		var zeroLevel []int
		for _, numberStr := range numbersStr {
			number, _ := strconv.Atoi(numberStr)
			zeroLevel = append(zeroLevel, number)
		}
		differences = append(differences, zeroLevel)

		for j := 0; ; j++ {
			allZeros := true
			var nextLevelDiferences []int
			for i := 0; i < len(differences[j])-1; i++ {
				diff := differences[j][i+1] - differences[j][i]
				nextLevelDiferences = append(nextLevelDiferences, diff)

				if diff != 0 {
					allZeros = false
				}
			}
			differences = append(differences, nextLevelDiferences)
			if allZeros {
				break
			}
		}

		differences[len(differences)-1] = append(differences[len(differences)-1], 0)

		for i := len(differences) - 2; i >= 0; i-- {
			nextLineLastItem := differences[i+1][len(differences[i+1])-1]
			currentLineLastItem := differences[i][len(differences[i])-1]
			differences[i] = append(differences[i], nextLineLastItem+currentLineLastItem)
		}

		res += differences[0][len(differences[0])-1]
	}

	return res
}

func puzzle2(lines []string) int {
	var res int

	for _, line := range lines {
		numbersStr := strings.Split(line, " ")

		var differences [][]int

		zeroLevel := make([]int, len(numbersStr))
		for i, numberStr := range numbersStr {
			number, _ := strconv.Atoi(numberStr)
			zeroLevel[len(numbersStr)-i-1] = number
		}
		differences = append(differences, zeroLevel)

		for j := 0; ; j++ {
			allZeros := true
			var nextLevelDiferences []int
			for i := 0; i < len(differences[j])-1; i++ {
				diff := differences[j][i+1] - differences[j][i]
				nextLevelDiferences = append(nextLevelDiferences, diff)

				if diff != 0 {
					allZeros = false
				}
			}
			differences = append(differences, nextLevelDiferences)
			if allZeros {
				break
			}
		}

		differences[len(differences)-1] = append(differences[len(differences)-1], 0)

		for i := len(differences) - 2; i >= 0; i-- {
			nextLineLastItem := differences[i+1][len(differences[i+1])-1]
			currentLineLastItem := differences[i][len(differences[i])-1]
			differences[i] = append(differences[i], nextLineLastItem+currentLineLastItem)
		}

		res += differences[0][len(differences[0])-1]
	}

	return res
}
