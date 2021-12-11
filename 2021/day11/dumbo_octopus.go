package main

import (
	"fmt"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day11/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	total := countTotalFlashes(data, 195)
	fmt.Println("total:", total)

	step := findWhenAllOctopusesFlash(data)
	fmt.Println("step:", step)
}

func countTotalFlashes(data []string, steps int) int {
	total := 0
	b := newBoard(data)
	for i := 0; i < steps; i++ {
		b.simulate()
		total += b.getJustFlashed()
	}
	return total
}

func findWhenAllOctopusesFlash(data []string) int {
	b := newBoard(data)
	step := 0
	for {
		b.simulate()
		step++
		if b.doAllOctopusesFlash() {
			return step
		}
	}
}
