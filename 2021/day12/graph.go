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

func (g *Graph) getDistinctPathsAmount() int {
	paths := g.findAllPaths("start", make(map[string]struct{}))
	return len(paths)
}

func (g *Graph) findAllPaths(current string, visited map[string]struct{}) []string {
	var pathsFromCurrent []string
	for _, next := range g.vertexes[current] {
		if next == "start" {
			continue
		}
		if next == "end" {
			pathsFromCurrent = append(pathsFromCurrent, "end")
			continue
		}
		if _, exists := visited[next]; exists {
			continue
		}

		if isCaveSmall(next) {
			visited[next] = struct{}{}
		}
		pathsFromNext := g.findAllPaths(next, visited)
		delete(visited, next)

		for _, fromNext := range pathsFromNext {
			pathsFromCurrent = append(pathsFromCurrent, strings.Join([]string{next, fromNext}, ","))
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
