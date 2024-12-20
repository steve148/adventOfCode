package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]int
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for _, r := range line {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				// Assume that conversion error due to non-numeric character.
				// Set to -1 to represent a unusable cell.
				num = -1
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	minX := 0
	maxX := len(grid[0]) - 1
	minY := 0
	maxY := len(grid) - 1

	var trailheads []Point
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 0 {
				trailheads = append(trailheads, Point{x, y})
			}
		}
	}

	fmt.Println(minX, maxX, minY, maxY)
	PrintGrid(grid)
	fmt.Println(trailheads)

	totalScore := 0
	for _, trailhead := range trailheads {
		queue := []Point{trailhead}
		reachable := make(map[Point]bool)

		for len(queue) > 0 {
			var p Point
			p, queue = queue[0], queue[1:]
			val := grid[p.y][p.x]

			if val == 9 {
				reachable[p] = true
				continue
			}

			if p.x < maxX && grid[p.y][p.x+1] == val+1 {
				queue = append(queue, Point{p.x + 1, p.y})
			}
			if p.x > minX && grid[p.y][p.x-1] == val+1 {
				queue = append(queue, Point{p.x - 1, p.y})
			}
			if p.y < maxY && grid[p.y+1][p.x] == val+1 {
				queue = append(queue, Point{p.x, p.y + 1})
			}
			if p.y > minY && grid[p.y-1][p.x] == val+1 {
				queue = append(queue, Point{p.x, p.y - 1})
			}
		}

		totalScore += len(reachable)
	}

	fmt.Println(totalScore)
}

func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
