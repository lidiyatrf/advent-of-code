package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ToStrings("2021/day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	straightOverlaps, err := getStraightOverlaps(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("straight overlaps:", straightOverlaps)

	allOverlaps, err := getAllOverlaps(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("all overlaps:", allOverlaps)
}

// getStraightOverlaps for part1
func getStraightOverlaps(data []string) (int, error) {
	plane := newPlane()
	for _, next := range data {
		plane.addStraightLine(next)
	}
	return plane.getOverlaps(), nil
}

// getAllOverlaps for part2
func getAllOverlaps(data []string) (int, error) {
	plane := newPlane()
	for _, next := range data {
		plane.addLine(next)
	}
	return plane.getOverlaps(), nil
}
