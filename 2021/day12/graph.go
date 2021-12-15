package main

import (
	"strings"
	"unicode"
)

const start = "start"
const end = "end"

type Graph struct {
	vertexes map[string][]string
}

func newGraph(data []string) *Graph {
	g := &Graph{
		vertexes: make(map[string][]string),
	}
	for _, next := range data {
		points := strings.Split(next, "-")
		g.vertexes[points[0]] = append(g.vertexes[points[0]], points[1])
		g.vertexes[points[1]] = append(g.vertexes[points[1]], points[0])
	}
	return g
}

func (g *Graph) getDistinctPathsAmountPart1() int {
	paths := g.findAllPaths(start, make(map[string]struct{}))
	return len(paths)
}

func (g *Graph) findAllPaths(currVertex string, visited map[string]struct{}) []string {
	var pathsFromCurrent []string
	for _, nextVertex := range g.vertexes[currVertex] {
		if nextVertex == start {
			continue
		}
		if nextVertex == end {
			pathsFromCurrent = append(pathsFromCurrent, end)
			continue
		}
		if _, exists := visited[nextVertex]; exists {
			continue
		}

		if isCaveSmall(nextVertex) {
			visited[nextVertex] = struct{}{}
		}
		pathsFromNext := g.findAllPaths(nextVertex, visited)
		delete(visited, nextVertex)

		for _, fromNext := range pathsFromNext {
			pathsFromCurrent = append(pathsFromCurrent, strings.Join([]string{nextVertex, fromNext}, ","))
		}
	}
	return pathsFromCurrent
}

func isCaveSmall(str string) bool {
	for _, letter := range str {
		if unicode.IsUpper(letter) {
			return false
		}
	}
	return true
}

func (g *Graph) getDistinctPathsAmountPart2() int {
	paths := g.findAllPathsPart2(start, make(map[string]struct{}), make(map[string]struct{}))
	return len(paths)
}

func (g *Graph) findAllPathsPart2(currVertex string, visited map[string]struct{}, visitedTwice map[string]struct{}) []string {
	var pathsFromCurrent []string

	for _, nextVertex := range g.vertexes[currVertex] {
		if nextVertex == start {
			continue
		}
		if nextVertex == end {
			pathsFromCurrent = append(pathsFromCurrent, end)
			continue
		}

		if _, exists := visited[nextVertex]; exists {
			if len(visitedTwice) == 1 {
				continue
			}
			visitedTwice[nextVertex] = struct{}{}
		} else if isCaveSmall(nextVertex) {
			visited[nextVertex] = struct{}{}
		}

		pathsFromNext := g.findAllPathsPart2(nextVertex, visited, visitedTwice)

		if _, exists := visitedTwice[nextVertex]; exists {
			delete(visitedTwice, nextVertex)
		} else {
			delete(visited, nextVertex)
		}

		for _, fromNext := range pathsFromNext {
			pathsFromCurrent = append(pathsFromCurrent, strings.Join([]string{nextVertex, fromNext}, ","))
		}
	}
	return pathsFromCurrent
}
