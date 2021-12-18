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
}

func getMagnitudeOfSum(data []string) int {
	sum := SnailfishNumber{value: data[0]}
	for _, next := range data[1:] {
		sum.add(next)
		sum.reduce()
	}
	return sum.getMagnitude()
}
