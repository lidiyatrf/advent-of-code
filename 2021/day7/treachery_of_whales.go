package main

import (
	"fmt"
	"math"
	"sort"

	"advent-of-code/2021/file"
	"advent-of-code/2021/parse"
)

func main() {
	data, err := file.ParseToStrings("2021/day7/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fuel, err := countFuelToAlignCrabs(data[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("requires: %d fuel", fuel)
}

func countFuelToAlignCrabs(data string) (fuel int, err error) {
	positions, err := parse.StringToInts(data, ",")
	if err != nil {
		return 0, err
	}

	median := getMedian(positions)
	fuel = countCostToAlign(positions, median)

	return fuel, nil
}

func getMedian(array []int) int {
	sort.Ints(array)
	if len(array)%2 == 1 {
		return array[len(array)/2+1]
	}
	return (array[len(array)/2-1] + array[len(array)/2]) / 2
}

func countCostToAlign(positions []int, mediana int) int {
	counter := 0
	for _, next := range positions {
		counter += int(math.Abs(float64(next - mediana)))
	}
	return counter
}
