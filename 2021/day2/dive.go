package main

import (
	"fmt"
	"strconv"
	"strings"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day2/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	result, err := move(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
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
