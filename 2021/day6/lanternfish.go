package main

import (
	"advent-of-code/file"
	"fmt"
)

func main() {
	data, err := file.ToStrings("2021/day6/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	lanternfish80, err := countLanternfish(data[0], 80)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("lanternfish after 80 days:", lanternfish80)

	lanternfish256, err := countLanternfish(data[0], 256)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("lanternfish after 256 days:", lanternfish256)
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
