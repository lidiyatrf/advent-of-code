package main

import (
	"fmt"
	"strings"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day8/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	amountOf1478, err := count1478digits(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("amount of digits 1 4 7 8:", amountOf1478)

	sum, err := sumOutputValues(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sum:", sum)
}

func count1478digits(data []string) (int, error) {
	counter := 0
	for _, nextLine := range data {
		splitted := strings.Split(nextLine, " | ")
		if len(splitted) != 2 {
			return 0, fmt.Errorf("error when splitting")
		}
		secondPartSplitted := strings.Split(splitted[1], " ")
		for _, nextDigit := range secondPartSplitted {
			if len(nextDigit) == 2 || len(nextDigit) == 3 || len(nextDigit) == 4 || len(nextDigit) == 7 {
				counter++
			}
		}
	}
	return counter, nil
}

func sumOutputValues(data []string) (int, error) {
	sum := 0
	for _, nextLine := range data {
		splitted := strings.Split(nextLine, " | ")
		if len(splitted) != 2 {
			return 0, fmt.Errorf("error when splitting by |")
		}

		signalWires := strings.Split(splitted[0], " ")
		if len(signalWires) != 10 {
			return 0, fmt.Errorf("error when splitting first part")
		}
		d := newDisplay(signalWires)

		secondPartSplitted := strings.Split(splitted[1], " ")
		if len(secondPartSplitted) != 4 {
			return 0, fmt.Errorf("error when splitting second part")
		}
		sum += d.decodeOutputValue(secondPartSplitted)
	}
	return sum, nil
}
