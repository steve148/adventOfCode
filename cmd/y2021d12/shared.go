package day12

import "strings"

type Graph struct {
	vertices map[string]bool
	edges    map[string][]string
}

func NewGraph() Graph {
	var g Graph
	g.vertices = make(map[string]bool)
	g.edges = make(map[string][]string)
	return g
}

func parseInput(lines []string) (g Graph) {
	g = NewGraph()
	for _, line := range lines {
		split := strings.Split(line, "-")
		l := split[0]
		r := split[1]

		g.vertices[l] = true
		g.vertices[r] = true
		g.edges[l] = append(g.edges[l], r)
		g.edges[r] = append(g.edges[r], l)
	}
	return
}

func isLower(s string) bool {
	return strings.ToLower(s) == s
}
