package main

import (
	"advent-of-code/file"
	"advent-of-code/utils"
	"fmt"
)

func main() {
	data, err := file.ParseToStrings("2023/day13/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
}

func puzzle1(lines []string) int {
	var res int

	puzzleStartLine := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			puzzleStartLine = i + 1
			continue
		}
		if i == len(lines)-1 || lines[i+1] == "" {
			rows, ok := findHorizontalMirror(lines[puzzleStartLine : i+1])
			if ok {
				res += 100 * rows
				puzzleStartLine = i + 1
				continue
			}
			columns, ok := findVerticalMirror(lines[puzzleStartLine : i+1])
			res += columns
			puzzleStartLine = i + 1
		}
	}

	return res
}

func findHorizontalMirror(lines []string) (int, bool) {
	for row := 0; row < len(lines)-1; row++ {
		if lines[row] != lines[row+1] {
			continue
		}

		rowsToCheck := utils.Min(row+1, len(lines)-row-1)

		mirror := true

		for i := 0; i < rowsToCheck; i++ {
			upRow := row - i
			downRow := row + i + 1

			if lines[upRow] != lines[downRow] {
				mirror = false
				break
			}
		}

		if !mirror {
			continue
		}

		return row + 1, true
	}

	return -1, false
}

func findVerticalMirror(lines []string) (int, bool) {
	for col := 0; col < len(lines[0])-1; col++ {
		mirror := true
		for row := 0; row < len(lines); row++ {
			if lines[row][col] != lines[row][col+1] {
				mirror = false
				break
			}
		}

		if !mirror {
			continue
		}

		columnsToCheck := utils.Min(col+1, len(lines[0])-col-1)

		for i := 0; i < columnsToCheck; i++ {
			leftColumn := col - i
			rightColumn := col + i + 1

			for row := 0; row < len(lines); row++ {
				if lines[row][leftColumn] != lines[row][rightColumn] {
					mirror = false
					break
				}
			}
		}

		if !mirror {
			continue
		}

		return col + 1, true
	}

	return -1, false
}
