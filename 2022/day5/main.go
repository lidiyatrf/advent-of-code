package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"advent-of-code/datastruct"
	"advent-of-code/file"
	"advent-of-code/split"
)

func main() {
	data, err := file.ToStrings("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle1:", puzzle1(data))
	fmt.Println("puzzle2:", puzzle2(data))
}

var movementRegexp = regexp.MustCompile(`^move (.*?) from (.*?) to (.*?)$`)

func puzzle1(lines []string) string {
	emptyLine := findEmptyLineIndex(lines)
	stacks := convertInputToStacks(lines[0:emptyLine])

	for _, line := range lines[emptyLine+1:] {
		howMany, from, to := extractMovements(line)

		for i := 0; i < howMany; i++ {
			err := stacks[from-1].MoveOneElementTo(&stacks[to-1])
			if err != nil {
				log.Fatalf("something went wrong: %v", err)
			}
		}
	}

	var result string
	for _, stack := range stacks {
		result += *stack.Peek()
	}
	return result
}

func puzzle2(lines []string) string {
	emptyLine := findEmptyLineIndex(lines)
	stacks := convertInputToStacks(lines[0:emptyLine])

	for _, line := range lines[emptyLine+1:] {
		howMany, from, to := extractMovements(line)

		if howMany == 1 {
			err := stacks[from-1].MoveOneElementTo(&stacks[to-1])
			if err != nil {
				log.Fatalf("something went wrong: %v", err)
			}
			continue
		}

		var extracted datastruct.Stack[string]
		for i := 0; i < howMany; i++ {
			err := stacks[from-1].MoveOneElementTo(&extracted)
			if err != nil {
				log.Fatalf("something went wrong: %v", err)
			}
		}
		for i := 0; i < howMany; i++ {
			err := extracted.MoveOneElementTo(&stacks[to-1])
			if err != nil {
				log.Fatalf("something went wrong: %v", err)
			}
		}
	}

	var result string
	for _, stack := range stacks {
		result += *stack.Peek()
	}
	return result
}

func findEmptyLineIndex(lines []string) int {
	for i, line := range lines {
		if line == "" {
			return i
		}
	}
	return -1
}

func extractMovements(line string) (int, int, int) {
	movements := movementRegexp.FindStringSubmatch(line)
	howMany, err := strconv.Atoi(movements[1])
	if err != nil {
		log.Fatal("cannot convert howMany to number")
	}
	from, err := strconv.Atoi(movements[2])
	if err != nil {
		log.Fatal("cannot convert from to number")
	}
	to, err := strconv.Atoi(movements[3])
	if err != nil {
		log.Fatal("cannot convert to to number")
	}
	return howMany, from, to
}

func convertInputToStacks(lines []string) []datastruct.Stack[string] {
	numbers, err := split.ToInts(lines[len(lines)-1], "   ")
	if err != nil {
		log.Fatal("cannot parse line with numbers")
	}

	stacks := make([]datastruct.Stack[string], len(numbers))
	for j := len(lines) - 2; j >= 0; j-- {
		currentLine := lines[j]
		for i := 0; i < len(stacks); i++ {
			symbol := currentLine[i*4+1 : i*4+2]
			if symbol != " " {
				stacks[i].Push(symbol)
			}
		}
	}
	return stacks
}
