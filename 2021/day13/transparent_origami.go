package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ToStrings("2021/day13/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	points := getPointsAfterFirstFold(data)
	fmt.Println("points:", points)

	printPointsAfterAllFolds(data)
}

func getPointsAfterFirstFold(data []string) int {
	points, folds := splitData(data)

	p := newPlane(points)
	p.makeFold(folds[0])
	return p.getPointsSize()
}

func splitData(data []string) ([]string, []string) {
	for i, next := range data {
		if next == "" {
			return data[:i], data[i+1:]
		}
	}
	return data, []string{}
}

func printPointsAfterAllFolds(data []string) {
	points, folds := splitData(data)
	p := newPlane(points)
	for _, nextFold := range folds {
		p.makeFold(nextFold)
	}
	p.printResult()
}
