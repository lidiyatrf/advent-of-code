package main

import (
	"advent-of-code/2021/file"
	"fmt"
)

func main() {
	data, err := file.ParseToStrings("2021/day9/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	sum, err := sumRiskedLevel(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("sum:", sum)
}

func sumRiskedLevel(data []string) (int, error) {
	hm, err := newHeightmap(data)
	if err != nil {
		return 0, nil
	}

	for i := 0; i < hm.getHeight(); i++ {
		for j := 0; j < hm.getWidth(); j++ {
			if hm.isSmallerThanAround(i, j) {
				hm.markLowest(i, j)
			}
		}
	}

	return hm.sumRiskedLevel(), nil
}
