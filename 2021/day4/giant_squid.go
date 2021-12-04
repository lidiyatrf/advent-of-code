package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day4/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	bestBoard, err := findBestBoard(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("best board result", bestBoard)

	worstBoard, err := findWorstBoard(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("best board result", worstBoard)
}

func findWorstBoard(data []string) (int, error) {
	numbers, err := parseNumbers(data[0])
	if err != nil {
		return 0, err
	}
	boards, err := parseBoards(data[1:])
	if err != nil {
		return 0, err
	}

	worstBoardIndex := 0
	minStepsLimit := 0
	for k := 0; k < len(boards); k++ {
		better, steps := getStepsToFail(numbers, &boards[k], minStepsLimit)
		if !better {
			continue
		}
		worstBoardIndex = k
		minStepsLimit = steps
	}
	return boards[worstBoardIndex].getSumOfUnmarked() * numbers[minStepsLimit], nil
}

func getStepsToFail(numbers []int, board *Board, minStepsLimit int) (failAfterLimit bool, steps int) {
	for k := 0; k < len(numbers); k++ {
		marked, i, j := board.Mark(numbers[k])
		if !marked {
			continue
		}
		if board.isRowCompleted(i) || board.isColumnCompleted(j) {
			return k > minStepsLimit, k
		}
	}
	return true, len(numbers)
}

func findBestBoard(data []string) (int, error) {
	numbers, err := parseNumbers(data[0])
	if err != nil {
		return 0, err
	}
	boards, err := parseBoards(data[1:])
	if err != nil {
		return 0, err
	}

	bestBoardIndex := 0
	maxStepsLimit := len(numbers)
	for k := 0; k < len(boards); k++ {
		better, steps := getStepsToWin(numbers, &boards[k], maxStepsLimit)
		if !better {
			continue
		}
		bestBoardIndex = k
		maxStepsLimit = steps
	}
	return boards[bestBoardIndex].getSumOfUnmarked() * numbers[maxStepsLimit], nil
}

func getStepsToWin(numbers []int, board *Board, maxStepsLimit int) (winBeforeLimit bool, steps int) {
	for k := 0; k < maxStepsLimit; k++ {
		marked, i, j := board.Mark(numbers[k])
		if !marked {
			continue
		}
		if board.isRowCompleted(i) || board.isColumnCompleted(j) {
			return true, k
		}
	}
	return false, 0
}

func parseNumbers(str string) ([]int, error) {
	var result []int
	strs := strings.Split(str, ",")
	for _, next := range strs {
		num, err := strconv.Atoi(next)
		if err != nil {
			return []int{}, fmt.Errorf("impossible to parse numbers")
		}
		result = append(result, num)
	}
	return result, nil
}
