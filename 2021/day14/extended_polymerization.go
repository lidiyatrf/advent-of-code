package main

import (
	"fmt"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day14/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	diff := getDiffMostAndLeastCommon(data, 10)
	fmt.Println("diff:", diff)
}

func getDiffMostAndLeastCommon(data []string, steps int) int {
	polymer, rules := splitData(data)
	p := newPolymer(polymer, rules)
	for i := 0; i < steps; i++ {
		p.polymerize()
	}
	return p.getMostCommonBitOccurrence() - p.getLeastCommonBitOccurrence()
}

func splitData(data []string) (string, []string) {
	return data[0], data[2:]
}
