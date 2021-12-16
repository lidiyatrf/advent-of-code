package main

import (
	"fmt"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day15/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	totalRisk := getLowestRiskPath(data) // 656
	fmt.Println("totalRisk", totalRisk)
}

func getLowestRiskPath(data []string) int {
	matrix := newMatrix(data)
	return matrix.getLowestRiskPath()
}
