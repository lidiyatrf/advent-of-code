package main

import (
	"advent-of-code/file"
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2023/day18/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

type point struct {
	i int
	j int
}

func puzzle1(lines []string) int {
	var pathLength int
	var corners []point
	startPoint := point{0, 0}

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		direction := splitted[0]
		length, _ := strconv.Atoi(splitted[1])

		pathLength += length

		switch direction {
		case "D":
			startPoint.i += length
		case "U":
			startPoint.i -= length
		case "R":
			startPoint.j += length
		case "L":
			startPoint.j -= length
		}
		corners = append(corners, startPoint)
	}

	return calculateSquare(corners) - pathLength/2 + 1 + pathLength
}

func calculateSquare(corners []point) int {
	var res int
	for i := 0; i < len(corners)-1; i++ {
		res += corners[i].i * corners[i+1].j
	}
	res += corners[len(corners)-1].i * corners[0].j

	for i := 0; i < len(corners)-1; i++ {
		res -= corners[i+1].i * corners[i].j
	}
	res -= corners[0].i * corners[len(corners)-1].j
	return utils.Abs(res / 2)
}

type pointInt64 struct {
	i int64
	j int64
}

func puzzle2(lines []string) int64 {
	var pathLength int64
	var corners []pointInt64
	startPoint := pointInt64{0, 0}

	for _, line := range lines {
		splitted := strings.Split(line, " ")
		hexPart := splitted[2]
		direction := hexPart[len(hexPart)-2 : len(hexPart)-1]
		length, _ := strconv.ParseInt(hexPart[2:len(hexPart)-2], 16, 64)

		pathLength += length

		switch direction {
		case "1":
			startPoint.i += length
		case "3":
			startPoint.i -= length
		case "0":
			startPoint.j += length
		case "2":
			startPoint.j -= length
		}
		corners = append(corners, startPoint)
	}

	return calculateSquareInt64(corners) - pathLength/2 + 1 + pathLength
}

func calculateSquareInt64(corners []pointInt64) int64 {
	var res int64
	for i := 0; i < len(corners)-1; i++ {
		res += corners[i].i * corners[i+1].j
	}
	res += corners[len(corners)-1].i * corners[0].j

	for i := 0; i < len(corners)-1; i++ {
		res -= corners[i+1].i * corners[i].j
	}
	res -= corners[0].i * corners[len(corners)-1].j

	if res < 0 {
		return res / 2 * (-1)
	}
	return res / 2
}
