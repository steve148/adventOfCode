package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func traverseSmall(start string, g Graph, alreadyVisited map[string]bool) int {
	fmt.Println(start, alreadyVisited)
	if start == "end" {
		return 1
	}
	visited := map[string]bool{start: strings.ToLower(start) == start}
	for vertex := range alreadyVisited {
		visited[vertex] = alreadyVisited[vertex]
	}

	var paths int
	for _, vertex := range g.edges[start] {
		if !visited[vertex] {
			paths += traverseSmall(vertex, g, visited)
		}
	}
	return paths
}

func Part1() {
	file, err := os.Open("./day12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	g := parseInput(lines)

	count := traverseSmall("start", g, map[string]bool{})
	fmt.Println(count)
}
