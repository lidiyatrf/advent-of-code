package main

import (
	"advent-of-code/file"
	"advent-of-code/parse"
	"fmt"
	"math"
	"sort"
)

func main() {
	data, err := file.ToStrings("2021/day7/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fuel1, err := countFuelToAlignCrabs1(data[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("part1 requires: %d fuel\n", fuel1)

	// brute force
	fuel2, err := countFuelToAlignCrabs2(data[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("part2 requires: %d fuel\n", fuel2)
}

func countFuelToAlignCrabs1(data string) (fuel int, err error) {
	positions, err := parse.StringToInts(data, ",")
	if err != nil {
		return 0, err
	}

	middle := getMedian(positions)
	fuel = calculateCost1(positions, middle)

	return fuel, nil
}

func getMedian(array []int) int {
	sort.Ints(array)
	if len(array)%2 == 1 {
		return array[len(array)/2+1]
	}
	return (array[len(array)/2-1] + array[len(array)/2]) / 2
}

func calculateCost1(positions []int, position int) int {
	counter := 0
	for _, next := range positions {
		counter += int(math.Abs(float64(next - position)))
	}
	return counter
}

func countFuelToAlignCrabs2(data string) (fuel int, err error) {
	positions, err := parse.StringToInts(data, ",")
	if err != nil {
		return 0, err
	}

	min := math.MaxInt
	for i := 0; i < 1200; i++ {
		result := calculateCost2(positions, i)
		if result < min {
			min = result
		}
	}
	return min, nil
}

func calculateCost2(array []int, position int) int {
	counter := 0
	for _, next := range array {
		delta := int(math.Abs(float64(next - position)))
		for i := 1; i <= delta; i++ {
			counter += i
		}
	}
	return counter
}
