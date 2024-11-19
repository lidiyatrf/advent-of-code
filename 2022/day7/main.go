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
	//fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(lines []string) int {
	var res int

	for _, line := range lines {
		res += len(line)

	}

	return res
}
