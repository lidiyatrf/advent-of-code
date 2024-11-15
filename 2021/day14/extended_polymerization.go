package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ToStrings("2021/day14/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	diff10 := getDiffMostAndLeastCommon(data, 10)
	fmt.Println("diff 10 steps:", diff10)

	diff40 := getDiffMostAndLeastCommon(data, 40)
	fmt.Println("diff 40 steps:", diff40)
}

func getDiffMostAndLeastCommon(data []string, steps int) int64 {
	polymer, rules := splitData(data)
	p := newPolymer(polymer, rules)
	for i := 0; i < steps; i++ {
		p.polymerize()
	}
	max, min := p.getCommonBitsOccurrence()
	return max - min
}

func splitData(data []string) (string, []string) {
	return data[0], data[2:]
}
