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
	//fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(lines []string) int {
	var res int

	for _, line := range lines {
		res += len(line)

	}

	return res
}
