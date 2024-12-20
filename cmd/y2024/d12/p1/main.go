package main

import (
	"bufio"
	"os"
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

	var grid [][]rune
	for scanner.Scan() {
		line := scanner.Text()
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		grid = append(grid, row)
	}

	minX := 0
	maxX := len(grid[0]) - 1
	minY := 0
	maxY := len(grid) - 1

	score := 0
	seen := make(map[Point]bool)
	for y, row := range grid {
		for x, r := range row {
			p := Point{x: x, y: y}

			if _, ok := seen[p]; ok {
				continue
			}

			var area int
			var perim int
			var curr Point
			queue := []Point{p}
			for len(queue) > 0 {
				curr, queue = queue[0], queue[1:]
				if _, ok := seen[curr]; ok {
					continue
				}
				seen[curr] = true
				area++

				if curr.x < maxX && grid[curr.y][curr.x+1] == r {
					queue = append(queue, Point{x: curr.x + 1, y: curr.y})
				} else {
					perim++
				}

				if curr.x > minX && grid[curr.y][curr.x-1] == r {
					queue = append(queue, Point{x: curr.x - 1, y: curr.y})
				} else {
					perim++
				}

				if curr.y < maxY && grid[curr.y+1][curr.x] == r {
					queue = append(queue, Point{x: curr.x, y: curr.y + 1})
				} else {
					perim++
				}

				if curr.y > minY && grid[curr.y-1][curr.x] == r {
					queue = append(queue, Point{x: curr.x, y: curr.y - 1})
				} else {
					perim++
				}
			}

			score += area * perim
		}
	}

	println(score)
}
