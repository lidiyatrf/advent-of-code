package main

import (
	"advent-of-code/2021/file"
	"fmt"
)

func main() {
	data, err := file.ParseToStrings("2021/day12/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	paths := getPathsCount(data)
	fmt.Println("paths:", paths)
}

func getPathsCount(data []string) int {
	g := newGraph(data)
	return g.getDistinctPathsAmount()
}
