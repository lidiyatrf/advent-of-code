package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ParseToStrings("2023/day16/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

type point struct {
	i int
	j int
}

var energized map[point]map[string]struct{}

func puzzle1(lines []string) int {
	energized = make(map[point]map[string]struct{})
	start := point{0, 0}
	pass(lines, start, "right")
	return len(energized)
}

func puzzle2(lines []string) int {
	max := 0
	for i := 0; i < len(lines[0]); i++ {
		energized = make(map[point]map[string]struct{})
		start := point{0, i}
		pass(lines, start, "down")
		if len(energized) > max {
			max = len(energized)
		}
	}

	for i := 0; i < len(lines[0]); i++ {
		energized = make(map[point]map[string]struct{})
		start := point{len(lines) - 1, i}
		pass(lines, start, "up")
		if len(energized) > max {
			max = len(energized)
		}
	}

	for i := 0; i < len(lines); i++ {
		energized = make(map[point]map[string]struct{})
		start := point{i, 0}
		pass(lines, start, "right")
		if len(energized) > max {
			max = len(energized)
		}
	}

	for i := 0; i < len(lines); i++ {
		energized = make(map[point]map[string]struct{})
		start := point{i, len(lines[0]) - 1}
		pass(lines, start, "left")
		if len(energized) > max {
			max = len(energized)
		}
	}

	return max
}

func pass(lines []string, current point, direction string) {
	for {
		if current.i >= len(lines) || current.j >= len(lines[0]) || current.i < 0 || current.j < 0 {
			return
		}

		if pointDirections, visited := energized[current]; visited {
			if _, visitedDirection := pointDirections[direction]; visitedDirection {
				return
			}
		} else {
			energized[current] = make(map[string]struct{})
			energized[current][direction] = struct{}{}
		}

		var next point
		switch lines[current.i][current.j] {
		case '.':
			switch direction {
			case "left":
				next.i = current.i
				next.j = current.j - 1
			case "right":
				next.i = current.i
				next.j = current.j + 1
			case "up":
				next.i = current.i - 1
				next.j = current.j
			case "down":
				next.i = current.i + 1
				next.j = current.j
			}
		case '|':
			switch direction {
			case "left", "right":
				up := point{current.i - 1, current.j}
				down := point{current.i + 1, current.j}
				pass(lines, up, "up")
				pass(lines, down, "down")
				return
			case "up":
				next.i = current.i - 1
				next.j = current.j
			case "down":
				next.i = current.i + 1
				next.j = current.j
			}
		case '-':
			switch direction {
			case "left":
				next.i = current.i
				next.j = current.j - 1
			case "right":
				next.i = current.i
				next.j = current.j + 1
			case "up", "down":
				left := point{current.i, current.j - 1}
				right := point{current.i, current.j + 1}
				pass(lines, left, "left")
				pass(lines, right, "right")
				return
			}
		case '/':
			switch direction {
			case "left":
				next.i = current.i + 1
				next.j = current.j
				direction = "down"
			case "right":
				next.i = current.i - 1
				next.j = current.j
				direction = "up"
			case "up":
				next.i = current.i
				next.j = current.j + 1
				direction = "right"
			case "down":
				next.i = current.i
				next.j = current.j - 1
				direction = "left"
			}
		case '\\':
			switch direction {
			case "left":
				next.i = current.i - 1
				next.j = current.j
				direction = "up"
			case "right":
				next.i = current.i + 1
				next.j = current.j
				direction = "down"
			case "up":
				next.i = current.i
				next.j = current.j - 1
				direction = "left"
			case "down":
				next.i = current.i
				next.j = current.j + 1
				direction = "right"
			}
		}
		current = next
	}
}
