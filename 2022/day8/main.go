package main

import (
	"fmt"
	"log"

	"advent-of-code/file"
)

func main() {
	data, err := file.ToStrings("./input.txt")
	if err != nil {
		log.Fatalf("cannot read file: %v", err)
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	// fmt.Println("puzzle 2:", puzzle2(data))
}

type position struct {
	row    int
	column int
}

func puzzle1(lines []string) int {
	visible := make(map[position]struct{})

	// horizontally: left to right; right to left
	for row, line := range lines[1 : len(lines)-1] {
		highestLeft := line[0]
		for i := 1; i < len(line)-1; i++ {
			if line[i] <= highestLeft {
				continue
			}
			visible[position{row: row + 1, column: i}] = struct{}{}
			if line[i] > highestLeft {
				highestLeft = line[i]
			}
		}

		highestRight := line[len(line)-1]
		for i := len(line) - 2; i > 0; i-- {
			if line[i] <= highestRight {
				continue
			}
			position := position{row: row + 1, column: i}
			if _, exists := visible[position]; !exists {
				visible[position] = struct{}{}
			}
			if line[i] > highestRight {
				highestRight = line[i]
			}
		}
	}

	// vertically top to bottom; left to right
	for column := 1; column < len(lines[0])-1; column++ {
		highestTop := lines[0][column]
		for row := 1; row < len(lines)-1; row++ {
			if lines[row][column] <= highestTop {
				continue
			}
			position := position{row: row, column: column}
			if _, exists := visible[position]; !exists {
				visible[position] = struct{}{}
			}
			if lines[row][column] > highestTop {
				highestTop = lines[row][column]
			}
		}
	}

	// vertically bottom to top; left to right
	for column := 1; column < len(lines[0])-1; column++ {
		highestBottom := lines[len(lines)-1][column]
		for row := len(lines) - 2; row > 0; row-- {
			if lines[row][column] <= highestBottom {
				continue
			}
			position := position{row: row, column: column}
			if _, exists := visible[position]; !exists {
				visible[position] = struct{}{}
			}
			if lines[row][column] > highestBottom {
				highestBottom = lines[row][column]
			}
		}
	}

	visibleEdges := len(lines)*2 + len(lines[0])*2 - 4
	return visibleEdges + len(visible)
}
