package main

import (
	"advent-of-code/file"
	"advent-of-code/utils"
	"fmt"
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

type point struct {
	i int
	j int
}

var corners []point

func puzzle2(lines []string) int {
	corners = nil

	s := findS(lines)
	path := goNext(lines, s, s)
	square := calculateSquare()

	return square - path/2 + 1
}

func calculateSquare() int {
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

func puzzle1(lines []string) int {
	s := findS(lines)
	pathLength := goNext(lines, s, s)
	return pathLength/2 + pathLength%2
}

func goNext(lines []string, currentPoint point, previousPoint point) int {
	currentValue := lines[currentPoint.i][currentPoint.j]

	var first, second point

	switch currentValue {
	case '|':
		first = point{i: currentPoint.i - 1, j: currentPoint.j}
		second = point{i: currentPoint.i + 1, j: currentPoint.j}
	case '-':
		first = point{i: currentPoint.i, j: currentPoint.j - 1}
		second = point{i: currentPoint.i, j: currentPoint.j + 1}
	case 'L':
		corners = append(corners, point{currentPoint.i, currentPoint.j})

		first = point{i: currentPoint.i - 1, j: currentPoint.j}
		second = point{i: currentPoint.i, j: currentPoint.j + 1}
	case 'J':
		corners = append(corners, point{currentPoint.i, currentPoint.j})

		first = point{i: currentPoint.i - 1, j: currentPoint.j}
		second = point{i: currentPoint.i, j: currentPoint.j - 1}
	case '7':
		corners = append(corners, point{currentPoint.i, currentPoint.j})

		first = point{i: currentPoint.i + 1, j: currentPoint.j}
		second = point{i: currentPoint.i, j: currentPoint.j - 1}
	case 'F':
		corners = append(corners, point{currentPoint.i, currentPoint.j})

		first = point{i: currentPoint.i + 1, j: currentPoint.j}
		second = point{i: currentPoint.i, j: currentPoint.j + 1}
	case 'S':
		if previousPoint != currentPoint {
			return 1
		}

		nextPoint := getNextFromS(lines, currentPoint)
		return goNext(lines, nextPoint, currentPoint)
	default:
		panic(fmt.Sprintf("cell [%d,%d] has unsupported value [%s]. previous point was [%d,%d]", currentPoint.i, currentPoint.j, string(lines[currentPoint.i][currentPoint.j]), previousPoint.i, previousPoint.j))
	}

	var nextPoint point
	if first != previousPoint {
		nextPoint = first
	} else {
		nextPoint = second
	}

	result := goNext(lines, nextPoint, currentPoint)
	return result + 1
}

func findS(lines []string) point {
	for i, line := range lines {
		if j := strings.Index(line, "S"); j != -1 {
			return point{i, j}
		}
	}
	panic("no S point")
}

func getNextFromS(lines []string, sPoint point) point {
	//up: |, 7, F
	if sPoint.i > 0 {
		p := point{i: sPoint.i - 1, j: sPoint.j}
		if lines[p.i][p.j] == '|' || lines[p.i][p.j] == '7' || lines[p.i][p.j] == 'F' {
			return p
		}
	}

	//down: |, L, J
	if sPoint.i < len(lines)-1 {
		p := point{i: sPoint.i + 1, j: sPoint.j}
		if lines[p.i][p.j] == '|' || lines[p.i][p.j] == 'L' || lines[p.i][p.j] == 'J' {
			return p
		}
	}

	//left: -, L
	if sPoint.j > 0 {
		p := point{i: sPoint.i, j: sPoint.j - 1}
		if lines[p.i][p.j] == '-' || lines[p.i][p.j] == 'L' {
			return p
		}
	}

	//right: -, 7
	if sPoint.j < len(lines[0])-1 {
		p := point{i: sPoint.i, j: sPoint.j + 1}
		if lines[p.i][p.j] == '-' || lines[p.i][p.j] == '7' {
			return p
		}
	}

	panic("no way from S")
}
