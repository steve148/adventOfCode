package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func traverseSmallTwice(start string, g Graph, alreadyVisited map[string]bool, smallVisited bool) int {
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
			paths += traverseSmallTwice(vertex, g, visited, smallVisited)
		} else if !smallVisited && vertex != "start" {
			paths += traverseSmallTwice(vertex, g, visited, true)
		}
	}
	return paths
}

func Part2() {
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

	count := traverseSmallTwice("start", g, map[string]bool{}, false)
	fmt.Println(count)
}
