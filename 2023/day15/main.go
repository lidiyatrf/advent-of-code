package main

import (
	"advent-of-code/file"
	"container/list"
	"fmt"
	"strings"
)

func main() {
	data, err := file.ReadLine("2023/day15/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("puzzle 1:", puzzle1(data))
	fmt.Println("puzzle 2:", puzzle2(data))
}

func puzzle1(line string) int {
	steps := strings.Split(line, ",")

	var sum int

	for _, step := range steps {
		var currentValue int
		for _, s := range step {
			currentValue += int(byte(s))
			currentValue *= 17
			currentValue %= 256
		}
		sum += currentValue
	}

	return sum
}

type lens struct {
	label       string
	focalLength int
}

func puzzle2(line string) int {
	steps := strings.Split(line, ",")

	boxes := make(map[int]*list.List)

	for _, step := range steps {
		var box int
		var label string
		var dash bool
		var equals bool
		var newFocalLength int

		for i, s := range step {
			if s == '-' {
				dash = true
				break
			}
			if s == '=' {
				equals = true
				newFocalLength = int(step[i+1] - '0')
				break
			}
			box += int(byte(s))
			box *= 17
			box %= 256

			label += string(s)
		}

		boxList, ok := boxes[box]

		if equals {
			if !ok {
				l := list.New()
				l.PushBack(&lens{label, newFocalLength})
				boxes[box] = l
				continue
			}

			var found bool
			for e := boxList.Front(); e != nil; e = e.Next() {
				l := e.Value.(*lens)
				if l.label == label {
					l.focalLength = newFocalLength
					found = true
					break
				}
			}

			if !found {
				boxList.PushBack(&lens{label, newFocalLength})
			}
			continue
		}

		if dash {
			if !ok {
				continue
			}

			for e := boxList.Front(); e != nil; e = e.Next() {
				l := e.Value.(*lens)
				if l.label == label {
					boxList.Remove(e)
					break
				}
			}
			continue
		}
	}

	var sum int
	for boxNum, l := range boxes {
		var i int
		for e := l.Front(); e != nil; e = e.Next() {
			i++
			sum += (boxNum + 1) * i * e.Value.(*lens).focalLength
		}
	}
	return sum
}
