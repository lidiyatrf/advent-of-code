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
	fmt.Println("worst board result", worstBoard)
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
	stepsLimit := len(numbers)
	for k := 0; k < len(boards); k++ {
		betterResult, newSteps := boards[k].countStepsToWinLessThanLimit(numbers, stepsLimit)
		if !betterResult {
			continue
		}
		bestBoardIndex = k
		stepsLimit = newSteps
	}
	return boards[bestBoardIndex].getSumOfUnmarked() * numbers[stepsLimit], nil
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
	stepsLimit := 0
	for k := 0; k < len(boards); k++ {
		betterResult, newSteps := boards[k].countStepsToWinMoreThanLimit(numbers, stepsLimit)
		if !betterResult {
			continue
		}
		worstBoardIndex = k
		stepsLimit = newSteps
	}
	return boards[worstBoardIndex].getSumOfUnmarked() * numbers[stepsLimit], nil
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

func parseBoards(data []string) ([]Board, error) {
	var result []Board

	for i := 0; i < len(data); {
		if data[i] == "" {
			i++
			continue
		}
		nextBoard, err := newBoard(data[i : i+5])
		i += 5
		if err != nil {
			return nil, err
		}
		result = append(result, *nextBoard)
	}

	return result, nil
}
