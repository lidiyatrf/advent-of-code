package main

import (
	"fmt"
	"math"
)

type Point struct {
	x int
	y int
}

type Diagram struct {
	diagram map[Point]int
}

func newDiagram() *Diagram {
	return &Diagram{diagram: map[Point]int{}}
}

func (d *Diagram) addStraightLine(point1, point2 Point) {
	if point1.x == point2.x { // vertical
		for y := min(point1.y, point2.y); y <= max(point1.y, point2.y); y++ {
			d.diagram[Point{x: point1.x, y: y}] += 1
		}
	}
	if point1.y == point2.y { // horizontal
		for x := min(point1.x, point2.x); x <= max(point1.x, point2.x); x++ {
			d.diagram[Point{x: x, y: point1.y}] += 1
		}
	}
}

// addLine for part 2
func (d *Diagram) addLine(point1, point2 Point) error {
	if point1.x != point2.x &&
		point1.y != point2.y &&
		abs(point1.y, point2.y) != abs(point1.x, point2.x) {
		return fmt.Errorf("line is neither straight, no 45' diagonal")
	}

	xDir := getDirection(point1.x, point2.x)
	yDir := getDirection(point1.y, point2.y)
	delta := max(abs(point1.y, point2.y), abs(point1.x, point2.x))

	for i, x, y := 0, point1.x, point1.y; i <= delta; i, x, y = i+1, x+xDir, y+yDir {
		d.diagram[Point{x: x, y: y}] += 1
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

func (d *Diagram) getOverlaps() int {
	counter := 0
	for _, v := range d.diagram {
		if v > 1 {
			counter++
		}
	}
	return counter
}
