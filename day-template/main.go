package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ParseToStrings("202*/day*/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
}

func puzzle1(arr []string) int {
	var res int

	for _, line := range arr {
		res += len(line)
	}

	return res
}
