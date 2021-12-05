package main

import (
	"advent-of-code/2021/file"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data, err := file.ParseToStrings("2021/day5/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	straightOverlapped, err := getStraightOverlappedAmount(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("straight overlapped:", straightOverlapped)

	overlapped, err := getOverlappedAmount(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("overlapped:", overlapped)
}

// getOverlappedAmount for part2
func getOverlappedAmount(data []string) (int, error) {
	d := newDiagram()
	for _, next := range data {
		point1, point2, err := parseRow(next)
		if err != nil {
			return 0, err
		}
		d.addLine(point1, point2)
	}
	return d.getOverlaps(), nil
}

// getStraightOverlappedAmount for part1
func getStraightOverlappedAmount(data []string) (int, error) {
	d := newDiagram()
	for _, next := range data {
		point1, point2, err := parseRow(next)
		if err != nil {
			return 0, err
		}
		d.addStraightLine(point1, point2)
	}
	return d.getOverlaps(), nil
}

func parseRow(str string) (Point, Point, error) {
	points := strings.Split(str, " -> ")
	if len(points) != 2 {
		return Point{}, Point{}, fmt.Errorf("cannot parse row %q", str)
	}
	point1, err := parsePoint(points[0])
	if err != nil {
		return Point{}, Point{}, fmt.Errorf("cannot parse point %q", points[0])
	}
	point2, err := parsePoint(points[1])
	if err != nil {
		return Point{}, Point{}, fmt.Errorf("cannot parse point %q", points[1])
	}
	return point1, point2, nil
}

func parsePoint(str string) (Point, error) {
	nums := strings.Split(str, ",")
	if len(nums) != 2 {
		return Point{}, fmt.Errorf("incorrect point format %q", str)
	}
	x, err := strconv.Atoi(nums[0])
	if err != nil {
		return Point{}, fmt.Errorf("cannot parse to int %q", nums[0])
	}
	y, err := strconv.Atoi(nums[1])
	if err != nil {
		return Point{}, fmt.Errorf("cannot parse to int %q", nums[1])
	}
	return Point{x: x, y: y}, nil
}
