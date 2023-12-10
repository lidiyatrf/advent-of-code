package main

import (
	"advent-of-code/file"
	"fmt"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2022/day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(arr []string) int {
	var score int
	for _, line := range arr {
		splitted := strings.Split(line, " ")
		opponentShape := getShape(splitted[0])
		myShape := getShape(splitted[1])
		score += (int(myShape) + 1) + winMatrix[myShape][opponentShape]
	}
	return score
}

type shape int

const (
	rock shape = iota
	paper
	scissors
)

var winMatrix = [][]int{
	{3, 0, 6},
	{6, 3, 0},
	{0, 6, 3},
}

func getShape(name string) shape {
	switch name {
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissors
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	default:
		panic("unsupported")
	}
}

func puzzle2(arr []string) int {
	var score int
	for _, line := range arr {
		splitted := strings.Split(line, " ")
		opponentShape := getShape(splitted[0])
		res := getExpectedResult(splitted[1])

		for i := 0; i < 3; i++ {
			if winMatrix[i][opponentShape] == res {
				score += res + (i + 1)
			}
		}
	}
	return score
}

// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win
func getExpectedResult(name string) int {
	switch name {
	case "X":
		return 0
	case "Y":
		return 3
	case "Z":
		return 6
	default:
		panic("unsupported")
	}
}
