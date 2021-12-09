package main

import (
	"fmt"
)

type Heightmap struct {
	heights      [][]int
	lowestPoints map[Point]int
}

type Point struct {
	i int
	j int
}

func newHeightmap(data []string) (*Heightmap, error) {
	var table [][]int
	for _, nextLine := range data {
		var row []int
		for _, nextVal := range nextLine {
			num := int(nextVal - '0')
			if num < 0 && num > 9 {
				return nil, fmt.Errorf("cannot parse not number")
			}
			row = append(row, num)
		}
		table = append(table, row)
	}

	return &Heightmap{
		heights:      table,
		lowestPoints: make(map[Point]int),
	}, nil
}

func (hm *Heightmap) getWidth() int {
	return len(hm.heights[0])
}

func (hm *Heightmap) getHeight() int {
	return len(hm.heights)
}

func (hm *Heightmap) isSmallerThanAround(i, j int) bool {
	left, right, up, down := true, true, true, true
	if i != 0 {
		up = hm.heights[i][j] < hm.heights[i-1][j]
	}
	if i != hm.getHeight()-1 {
		down = hm.heights[i][j] < hm.heights[i+1][j]
	}
	if j != 0 {
		left = hm.heights[i][j] < hm.heights[i][j-1]
	}
	if j != hm.getWidth()-1 {
		right = hm.heights[i][j] < hm.heights[i][j+1]
	}
	return left && right && up && down
}

func (hm *Heightmap) markLowest(i, j int) {
	hm.lowestPoints[Point{i, j}] = hm.heights[i][j]

	delete(hm.lowestPoints, Point{i - 1, j})
	delete(hm.lowestPoints, Point{i + 1, j})
	delete(hm.lowestPoints, Point{i, j - 1})
	delete(hm.lowestPoints, Point{i, j + 1})
}

func (hm *Heightmap) sumRiskedLevel() int {
	sum := 0
	for _, v := range hm.lowestPoints {
		sum = sum + v + 1
	}
	return sum
}

func (hm *Heightmap) get3LargestBasins() []int {
	sizes := []int{0, 0, 0}

	for k, _ := range hm.lowestPoints {
		newSize := hm.getBasinSize(k)
		updateTop3(sizes, newSize)
	}

	return sizes
}

func updateTop3(array []int, newInt int) {
	if newInt > array[0] {
		array[2] = array[1]
		array[1] = array[0]
		array[0] = newInt
		return
	}
	if newInt > array[1] {
		array[2] = array[1]
		array[1] = newInt
		return
	}
	if newInt > array[2] {
		array[2] = newInt
		return
	}
}

func (hm *Heightmap) getBasinSize(basinLowestPoint Point) int {
	marked := make(map[Point]struct{})
	marked[basinLowestPoint] = struct{}{}
	hm.populateBasinPoints(basinLowestPoint, marked)
	return len(marked)
}

func (hm *Heightmap) populateBasinPoints(basinPoint Point, marked map[Point]struct{}) {
	i := basinPoint.i
	j := basinPoint.j

	if i != 0 {
		newPoint := Point{i - 1, j}
		if hm.doesPointBelongToBasin(basinPoint, newPoint) {
			marked[newPoint] = struct{}{}
			hm.populateBasinPoints(newPoint, marked)
		}
	}
	if i != hm.getHeight()-1 {
		newPoint := Point{i + 1, j}
		if hm.doesPointBelongToBasin(basinPoint, newPoint) {
			marked[newPoint] = struct{}{}
			hm.populateBasinPoints(newPoint, marked)
		}
	}
	if j != 0 {
		newPoint := Point{i, j - 1}
		if hm.doesPointBelongToBasin(basinPoint, newPoint) {
			marked[newPoint] = struct{}{}
			hm.populateBasinPoints(newPoint, marked)
		}
	}
	if j != hm.getWidth()-1 {
		newPoint := Point{i, j + 1}
		if hm.doesPointBelongToBasin(basinPoint, newPoint) {
			marked[newPoint] = struct{}{}
			hm.populateBasinPoints(newPoint, marked)
		}
	}
}

func (hm *Heightmap) doesPointBelongToBasin(basinPoint, newPoint Point) bool {
	return hm.heights[newPoint.i][newPoint.j] != 9 &&
		hm.heights[newPoint.i][newPoint.j] > hm.heights[basinPoint.i][basinPoint.j]
}
