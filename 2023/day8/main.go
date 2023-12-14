package main

import (
	"advent-of-code/file"
	"fmt"
	"regexp"
)

func main() {
	data, err := file.ParseToStrings("2023/day8/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

type leftRightInstructions struct {
	left  string
	right string
}

var instructionsRegexp = regexp.MustCompile("[A-Z0-9]{3}")

func puzzle1(lines []string) int {
	directions := lines[0]
	rules := make(map[string]leftRightInstructions)

	for _, line := range lines[2:] {
		numbers := instructionsRegexp.FindAllString(line, 3)
		rules[numbers[0]] = leftRightInstructions{numbers[1], numbers[2]}
	}

	current := "AAA"

	for i := 0; ; i++ {
		d := directions[i%len(directions)]
		next := rules[current]

		switch d {
		case 'L':
			current = next.left
		case 'R':
			current = next.right
		}

		if current == "ZZZ" {
			return i + 1
		}
	}
}

func puzzle2(lines []string) int {
	directions := lines[0]
	rules := make(map[string]leftRightInstructions)

	var starts []string
	for _, line := range lines[2:] {
		numbers := instructionsRegexp.FindAllString(line, 3)
		rules[numbers[0]] = leftRightInstructions{numbers[1], numbers[2]}

		if numbers[0][2] == 'A' {
			starts = append(starts, numbers[0])
		}
	}

	var minimumSteps []int

	for _, current := range starts {
		for i := 0; ; i++ {
			d := directions[i%len(directions)]
			next := rules[current]

			switch d {
			case 'L':
				current = next.left
			case 'R':
				current = next.right
			}

			if current[2] == 'Z' {
				minimumSteps = append(minimumSteps, i+1)
				break
			}
		}
	}

	return LCM(minimumSteps[0], minimumSteps[1], minimumSteps[2:]...)
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
