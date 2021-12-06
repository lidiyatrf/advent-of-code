package main

import (
	"fmt"

	"advent-of-code/2021/file"
)

func main() {
	data, err := file.ParseToStrings("2021/day6/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lanternfish, err := countLanternfish(data[0], 80)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("lanternfish:", lanternfish)
}

func countLanternfish(data string, days int) (int, error) {
	s, err := newSchool(data)
	if err != nil {
		return 0, err
	}
	for i := 0; i < days; i++ {
		s.grow()
	}
	return s.getTotal(), nil
}
