package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ToStrings("2021/day12/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	paths1 := getPathsCountPart1(data)
	fmt.Println("paths1:", paths1)

	paths2 := getPathsCountPart2(data)
	fmt.Println("paths2:", paths2)
}

func getPathsCountPart1(data []string) int {
	g := newGraph(data)
	return g.getDistinctPathsAmountPart1()
}

func getPathsCountPart2(data []string) int {
	g := newGraph(data)
	return g.getDistinctPathsAmountPart2()
}
