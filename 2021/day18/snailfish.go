package main

import (
	"fmt"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day18/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	magnitude := getMagnitudeOfSum(data)
	fmt.Println("magnitude", magnitude)

	largestMagnitude := getLargestMagnitudeOfSum(data)
	fmt.Println("largestMagnitude", largestMagnitude)
}

func getLargestMagnitudeOfSum(data []string) interface{} {
	var largestMagnitude int
	for i, first := range data {
		for j, second := range data {
			if i == j {
				continue
			}
			currentMagnitude := getMagnitudeOfSum([]string{first, second})
			if currentMagnitude > largestMagnitude {
				largestMagnitude = currentMagnitude
			}
		}
	}
	return largestMagnitude
}

func getMagnitudeOfSum(data []string) int {
	sum := SnailfishNumber{value: data[0]}
	for _, next := range data[1:] {
		sum.add(next)
		sum.reduce()
	}
	return sum.getMagnitude()
}
