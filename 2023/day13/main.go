package main

import (
	"advent-of-code/file"
	math1 "advent-of-code/utils"
	"fmt"
	"math"
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

		rowsToCheck := math1.Min(row+1, len(lines)-row-1)

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

		columnsToCheck := math1.Min(col+1, len(lines[0])-col-1)

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

func puzzle2(lines []string) int {
	var res int

	puzzleStartLine := 0

	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			puzzleStartLine = i + 1
			continue
		}
		if i == len(lines)-1 || lines[i+1] == "" {
			rows, ok := findHorizontalMirror2(lines[puzzleStartLine : i+1])
			if ok {
				res += 100 * rows
				puzzleStartLine = i + 1
				continue
			}
			columns, ok := findVerticalMirror2(lines[puzzleStartLine : i+1])
			res += columns
			puzzleStartLine = i + 1
		}
	}

	return res
}

func findHorizontalMirror2(lines []string) (int, bool) {
	var nums []uint64
	for _, line := range lines {
		var num uint64
		for i, r := range line {
			switch r {
			case '#':
				num += uint64(math.Pow(2, float64(i)))
			case '.':
				continue
			}
		}
		nums = append(nums, num)
	}

	for row := 0; row < len(nums)-1; row++ {
		res := nums[row] ^ nums[row+1]
		if res != 0 && res&(res-1) != 0 {
			continue
		}

		rowsToCheck := math1.Min(row+1, len(nums)-row-1)

		mirror := true
		smudge := false

		for i := 0; i < rowsToCheck; i++ {
			upRow := row - i
			downRow := row + i + 1

			res := nums[upRow] ^ nums[downRow]
			if res == 0 {
				continue
			}

			if !smudge && res&(res-1) == 0 {
				smudge = true
				continue
			}

			mirror = false
			smudge = false
			break
		}

		if !mirror || !smudge {
			continue
		}

		return row + 1, true
	}

	return 0, false
}

func findVerticalMirror2(lines []string) (int, bool) {
	var nums []uint64
	for j := 0; j < len(lines[0]); j++ {
		var num uint64
		for i := range lines {
			r := lines[i][j]
			switch r {
			case '#':
				num += uint64(math.Pow(2, float64(i)))
			case '.':
				continue
			}
		}
		nums = append(nums, num)
	}

	for row := 0; row < len(nums)-1; row++ {
		res := nums[row] ^ nums[row+1]
		if res != 0 && res&(res-1) != 0 {
			continue
		}

		rowsToCheck := math1.Min(row+1, len(nums)-row-1)

		mirror := true
		smudge := false

		for i := 0; i < rowsToCheck; i++ {
			upRow := row - i
			downRow := row + i + 1

			res := nums[upRow] ^ nums[downRow]
			if res == 0 {
				continue
			}

			if !smudge && res&(res-1) == 0 {
				smudge = true
				continue
			}

			mirror = false
			smudge = false
			break
		}

		if !mirror || !smudge {
			continue
		}

		return row + 1, true
	}

	return 0, false
}
