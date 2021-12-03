package main

import (
	"fmt"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToInts("2021/day1/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(countIncreasedTriple(data))
}

func countIncreasedTriple(arr []int) int {
	if len(arr) < 4 {
		return 0
	}

	prev := arr[0] + arr[1] + arr[2]
	counter := 0

	for i, _ := range arr[3:] {
		curr := prev + arr[i+3] - arr[i]
		if curr > prev {
			counter++
		}
		prev = curr
	}
	return counter
}

func countIncreasedSingle(arr []int) int {
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
