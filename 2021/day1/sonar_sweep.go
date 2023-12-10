package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ParseToInts("2021/day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("increased singles:", countIncreasedSingles(data))

	fmt.Println("increased triples:", countIncreasedTriples(data))
}

func countIncreasedSingles(arr []int) int {
	if len(arr) == 1 {
		return 0
	}

	prev := arr[0]
	counter := 0

	for _, next := range arr[1:] {
		if next > prev {
			counter++
		}
		prev = next
	}
	return counter
}

func countIncreasedTriples(arr []int) int {
	if len(arr) < 4 {
		return 0
	}

	prev := arr[0] + arr[1] + arr[2]
	counter := 0

	for i := 3; i < len(arr); i++ {
		curr := prev + arr[i] - arr[i-3]
		if curr > prev {
			counter++
		}
		prev = curr
	}
	return counter
}
