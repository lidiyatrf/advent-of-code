package main

import (
	"advent-of-code/file"
	"advent-of-code/utils"
	"fmt"
)

func main() {
	data, err := file.ParseToStrings("2023/day11/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

type point struct {
	row    int64
	column int64
}

func puzzle1(lines []string) int64 {
	return puzzle(lines, 2)
}

func puzzle2(lines []string) int64 {
	return puzzle(lines, 1000000)
}

func puzzle(lines []string, replacementForEmptyRowsAndColumns int64) int64 {
	var galaxies []point

	columnsContainGalaxies := make([]bool, len(lines[0]))
	var galaxyRowNumber int64

	for _, row := range lines {
		var rowContainsGalaxies bool
		for columnIndex, cellValue := range row {
			if cellValue == '#' {
				rowContainsGalaxies = true
				columnsContainGalaxies[columnIndex] = true

				galaxies = append(galaxies, point{int64(galaxyRowNumber), int64(columnIndex)})
			}
		}
		galaxyRowNumber++
		if !rowContainsGalaxies {
			galaxyRowNumber += replacementForEmptyRowsAndColumns - 1
		}
	}

	for j := len(columnsContainGalaxies) - 1; j >= 0; j-- {
		if columnsContainGalaxies[j] {
			continue
		}
		for i := 0; i < len(galaxies); i++ {
			if galaxies[i].column > int64(j) {
				galaxies[i].column += replacementForEmptyRowsAndColumns - 1
			}
		}
	}

	var res int64
	for i := 0; i < len(galaxies)-1; i++ {
		firstGalaxy := galaxies[i]
		for _, secondGalaxy := range galaxies[i+1:] {
			res += utils.Abs(secondGalaxy.column-firstGalaxy.column) + utils.Abs(secondGalaxy.row-firstGalaxy.row)
		}
	}

	return res
}
