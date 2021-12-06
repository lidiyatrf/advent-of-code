package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Plane struct {
	plane map[Point]int
}

func newPlane() *Plane {
	return &Plane{plane: map[Point]int{}}
}

// addStraightLine adds horizontal and vertical lines
func (p *Plane) addStraightLine(coordinates string) error {
	point1, point2, err := parseCoordinates(coordinates)
	if err != nil {
		return err
	}

	if point1.x == point2.x { // vertical
		for y := min(point1.y, point2.y); y <= max(point1.y, point2.y); y++ {
			p.plane[Point{x: point1.x, y: y}] += 1
		}
		return nil
	}
	if point1.y == point2.y { // horizontal
		for x := min(point1.x, point2.x); x <= max(point1.x, point2.x); x++ {
			p.plane[Point{x: x, y: point1.y}] += 1
		}
		return nil
	}
	// skipping diagonal lines
	return nil
}

// addLine adds horizontal, vertical and diagonal (45 degrees) lines
func (p *Plane) addLine(coordinates string) error {
	point1, point2, err := parseCoordinates(coordinates)
	if err != nil {
		return err
	}

	if point1.x != point2.x &&
		point1.y != point2.y &&
		abs(point1.y, point2.y) != abs(point1.x, point2.x) {
		return fmt.Errorf("line is neither straight, no 45' diagonal")
	}

	xDir := getDirection(point1.x, point2.x)
	yDir := getDirection(point1.y, point2.y)
	delta := max(abs(point1.y, point2.y), abs(point1.x, point2.x))

	for i, x, y := 0, point1.x, point1.y; i <= delta; i, x, y = i+1, x+xDir, y+yDir {
		p.plane[Point{x: x, y: y}] += 1
	}
	return nil
}

func getDirection(c1, c2 int) int {
	if c2 > c1 {
		return 1
	}
	if c1 > c2 {
		return -1
	}
	return 0
}

func min(n1, n2 int) int {
	return int(math.Min(float64(n1), float64(n2)))
}

func max(n1, n2 int) int {
	return int(math.Max(float64(n1), float64(n2)))
}

func abs(n1, n2 int) int {
	return int(math.Abs(float64(n1 - n2)))
}

func (p *Plane) getOverlaps() int {
	counter := 0
	for _, v := range p.plane {
		if v > 1 {
			counter++
		}
	}
	return counter
}

func parseCoordinates(str string) (Point, Point, error) {
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
