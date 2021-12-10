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

	sum, err := count1478digits(data)
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
