package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

type Coordinate struct {
	x, y int
}

func neighbor_coordinates(c Coordinate, width, height int) []Coordinate {
	neighbors := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var result []Coordinate
	for _, neighbor := range neighbors {
		new_x := c.x + neighbor[0]
		new_y := c.y + neighbor[1]

		if new_x >= 0 && new_x < width && new_y >= 0 && new_y < height {
			result = append(result, Coordinate{x: new_x, y: new_y})
		}
	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grid := make([][]rune, 0)
	var starts []Coordinate
	var end Coordinate

	for i, line := range lines {
		var row []rune
		for j, r := range line {
			if r == 'S' || r == 'a' {
				starts = append(starts, Coordinate{x: j, y: i})
				r = 'a'
			}
			if r == 'E' {
				end = Coordinate{x: j, y: i}
				r = 'z'
			}
			row = append(row, r)
		}
		grid = append(grid, row)
	}

	width := len(grid[0])
	height := len(grid)

	shortest := math.MaxInt
	for _, start := range starts {
		visited := make(map[Coordinate]bool)
		queue := []Coordinate{start}
		distance := map[Coordinate]int{start: 0}

		for len(queue) != 0 {
			// Pop coordinate off of queue.
			current := queue[0]
			queue = queue[1:]

			visited[current] = true

			if current == end {
				if distance[end] < shortest {
					shortest = distance[end]
				}
				break
			}

			for _, neighbor := range neighbor_coordinates(current, width, height) {
				if !visited[neighbor] && grid[neighbor.y][neighbor.x]-grid[current.y][current.x] <= 1 {
					if distance[neighbor] == 0 {
						queue = append(queue, neighbor)
						distance[neighbor] = distance[current] + 1
					}
					if distance[neighbor] >= distance[current]+1 {
						distance[neighbor] = distance[current] + 1
					}
				}
			}

			sort.Slice(queue, func(i, j int) bool {
				return distance[queue[i]] < distance[queue[j]]
			})
		}
	}

	fmt.Println(shortest)
}
