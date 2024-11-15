package main

import "advent-of-code/split"

type Point struct {
	i int
	j int
}

type RiskLevelMap struct {
	values [][]int
}

func newMatrix(lines []string) *RiskLevelMap {
	var matrix [][]int

	for _, line := range lines {
		ints, err := split.ToInts(line, "")
		if err != nil {
			panic(err)
		}
		matrix = append(matrix, ints)
	}

	return &RiskLevelMap{
		values: matrix,
	}
}

func (r *RiskLevelMap) getLowestRiskPath() int {
	var nextToVisit []Point
	totalRisk := make(map[Point]int)
	visited := make(map[Point]struct{})

	nextToVisit = append(nextToVisit, Point{0, 0})
	totalRisk[Point{0, 0}] = 0

	for len(nextToVisit) != 0 {
		current := nextToVisit[0]
		nextToVisit = nextToVisit[1:]
		if _, wasVisited := visited[current]; wasVisited {
			continue
		}

		nextAround := r.getNext(current)
		currentTotalRisk := totalRisk[current]
		visited[current] = struct{}{}

		for _, point := range nextAround {
			nextToVisit = append(nextToVisit, point)

			pointTotalRisk := totalRisk[point]
			newPointTotalRisk := currentTotalRisk + r.getRisk(point)
			if pointTotalRisk == 0 || pointTotalRisk > newPointTotalRisk {
				totalRisk[point] = newPointTotalRisk
				totalRisk = r.updateLeft(point, totalRisk)
			}
		}
	}

	return totalRisk[Point{r.getHeight() - 1, r.getWidth() - 1}]
}

func (r *RiskLevelMap) getNext(current Point) []Point {
	j := current.j
	i := current.i

	var result []Point

	if j != r.getWidth()-1 {
		result = append(result, Point{i: i, j: j + 1})
	}
	if i != r.getHeight()-1 {
		result = append(result, Point{i: i + 1, j: j})
	}
	return result
}

func (r *RiskLevelMap) updateLeft(current Point, totalRisk map[Point]int) map[Point]int {
	i := current.i
	j := current.j

	if j != 0 {
		if totalRisk[Point{i, j - 1}] > totalRisk[Point{i, j}]+r.getRisk(current) {
			totalRisk[Point{i, j - 1}] = totalRisk[Point{i, j}] + r.getRisk(current)
		}
	}

	return totalRisk
}

func (r *RiskLevelMap) getRisk(p Point) int {
	return r.values[p.i][p.j]
}

func (r *RiskLevelMap) getHeight() int {
	return len(r.values)
}

func (r *RiskLevelMap) getWidth() int {
	return len(r.values[0])
}
