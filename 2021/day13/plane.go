package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var foldRegex = regexp.MustCompile("^fold along (.*?)=(.*?)$")

type Point struct {
	x int
	y int
}

type Plane struct {
	points map[Point]struct{}
}

func newPoint(point string) Point {
	coordinates := strings.Split(point, ",")
	if len(coordinates) != 2 {
		panic("coordinates have wrong length")
	}
	x, err := strconv.Atoi(coordinates[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		panic(err)
	}
	return Point{x: x, y: y}
}

func newPlane(points []string) *Plane {
	plane := &Plane{
		points: make(map[Point]struct{}),
	}
	for _, next := range points {
		point := newPoint(next)
		plane.appendPoint(point)
	}
	return plane
}

func (p Plane) appendPoint(point Point) {
	p.points[point] = struct{}{}
}

func (p *Plane) makeFold(fold string) {
	axis, coordinate := parseFold(fold)

	if axis == 'x' {
		p.foldLeft(coordinate)
	}
	if axis == 'y' {
		p.foldUp(coordinate)
	}
}

func parseFold(fold string) (rune, int) {
	match := foldRegex.FindStringSubmatch(fold)
	amount, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}
	return []rune(match[1])[0], amount
}

func (p *Plane) foldLeft(x int) {
	for current, _ := range p.points {
		if current.x > x {
			delete(p.points, current)
			newX := 2*x - current.x
			p.points[Point{x: newX, y: current.y}] = struct{}{}
		}
	}
}

func (p *Plane) foldUp(y int) {
	for current, _ := range p.points {
		if current.y > y {
			delete(p.points, current)
			newY := 2*y - current.y
			p.points[Point{x: current.x, y: newY}] = struct{}{}
		}
	}
}

func (p Plane) getPointsSize() int {
	return len(p.points)
}

func (p Plane) printResult() {
	maxX := p.getMaxX()
	maxY := p.getMaxY()

	for j := 0; j < maxY; j++ {
		for i := 0; i < maxX; i++ {
			if _, exists := p.points[Point{x: i, y: j}]; exists {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (p Plane) getMaxX() int {
	maxX := 0
	for k, _ := range p.points {
		if k.x > maxX {
			maxX = k.x
		}
	}
	return maxX + 1
}

func (p Plane) getMaxY() int {
	maxY := 0
	for k, _ := range p.points {
		if k.y > maxY {
			maxY = k.y
		}
	}
	return maxY + 1
}
