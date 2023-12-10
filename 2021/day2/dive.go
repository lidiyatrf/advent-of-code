package main

import (
	"advent-of-code/file"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2021/day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	moveResult, err := move(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("move:", moveResult)

	moveWithAimResult, err := moveWithAim(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("move with aim:", moveWithAimResult)
}

func move(actions []string) (int, error) {
	initX, initY := 0, 0

	for _, next := range actions {
		x, y, err := getShift(next)
		if err != nil {
			return 0, err
		}
		initX += x
		initY += y
	}

	return initX * initY, nil
}

func getShift(action string) (x, y int, err error) {
	splitted := strings.Split(action, " ")

	if len(splitted) != 2 {
		return 0, 0, fmt.Errorf("invalid data")
	}

	direction := splitted[0]
	amount, err := strconv.Atoi(splitted[1])
	if err != nil {
		return 0, 0, err
	}

	switch direction {
	case "forward":
		return amount, 0, nil
	case "down":
		return 0, amount, nil
	case "up":
		return 0, -amount, nil
	default:
		return 0, 0, fmt.Errorf("unsupported direction")
	}
}

func moveWithAim(actions []string) (int, error) {
	initX, initY, initAim := 0, 0, 0

	for _, next := range actions {
		x, y, aim, err := getShiftWithAim(next, initAim)
		if err != nil {
			return 0, err
		}
		initX += x
		initY += y
		initAim += aim
	}

	return initX * initY, nil
}

func getShiftWithAim(action string, currAim int) (x, y, aim int, err error) {
	splitted := strings.Split(action, " ")

	if len(splitted) != 2 {
		return 0, 0, 0, fmt.Errorf("invalid data")
	}

	direction := splitted[0]
	amount, err := strconv.Atoi(splitted[1])
	if err != nil {
		return 0, 0, 0, err
	}

	switch direction {
	case "forward":
		return amount, amount * currAim, 0, nil
	case "down":
		return 0, 0, amount, nil
	case "up":
		return 0, 0, -amount, nil
	default:
		return 0, 0, 0, fmt.Errorf("unsupported direction")
	}
}
