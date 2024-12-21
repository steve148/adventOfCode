package main

import (
	"bufio"
	"fmt"
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
			var sides int
			var curr Point
			queue := []Point{p}
			for len(queue) > 0 {
				curr, queue = queue[0], queue[1:]
				if _, ok := seen[curr]; ok {
					continue
				}
				seen[curr] = true
				area++

				// Queue neighbours which are part of the region.
				if curr.x < maxX && grid[curr.y][curr.x+1] == r {
					queue = append(queue, Point{x: curr.x + 1, y: curr.y})
				}
				if curr.x > minX && grid[curr.y][curr.x-1] == r {
					queue = append(queue, Point{x: curr.x - 1, y: curr.y})
				}
				if curr.y < maxY && grid[curr.y+1][curr.x] == r {
					queue = append(queue, Point{x: curr.x, y: curr.y + 1})
				}
				if curr.y > minY && grid[curr.y-1][curr.x] == r {
					queue = append(queue, Point{x: curr.x, y: curr.y - 1})
				}

				left := Point{x: curr.x - 1, y: curr.y}
				up := Point{x: curr.x, y: curr.y - 1}
				right := Point{x: curr.x + 1, y: curr.y}
				down := Point{x: curr.x, y: curr.y + 1}

				noLeft := left.x == -1
				noUp := up.y == -1
				noRight := right.x == len(grid[0])
				noDown := down.y == len(grid)

				fmt.Println(curr)

				// Top left corner.
				if (noLeft && noUp) || (noLeft && !noUp && grid[curr.y-1][curr.x] != r) || (noUp && !noLeft && grid[curr.y][curr.x-1] != r) || (!noLeft && !noUp && grid[curr.y][curr.x-1] != r && grid[curr.y-1][curr.x] != r) || (!noLeft && !noUp && grid[curr.y][curr.x-1] == r && grid[curr.y-1][curr.x] == r && grid[curr.y-1][curr.x-1] != r) {
					sides++
				}
				// Top right corner.
				if (noRight && noUp) || (noRight && !noUp && grid[curr.y-1][curr.x] != r) || (noUp && !noRight && grid[curr.y][curr.x+1] != r) || (!noRight && !noUp && grid[curr.y][curr.x+1] != r && grid[curr.y-1][curr.x] != r) || (!noRight && !noUp && grid[curr.y][curr.x+1] == r && grid[curr.y-1][curr.x] == r && grid[curr.y-1][curr.x+1] != r) {
					sides++
				}
				// Bottom left corner.
				if (noLeft && noDown) || (noLeft && !noDown && grid[curr.y+1][curr.x] != r) || (noDown && !noLeft && grid[curr.y][curr.x-1] != r) || (!noLeft && !noDown && grid[curr.y][curr.x-1] != r && grid[curr.y+1][curr.x] != r) || (!noLeft && !noDown && grid[curr.y][curr.x-1] == r && grid[curr.y+1][curr.x] == r && grid[curr.y+1][curr.x-1] != r) {
					sides++
				}
				// Bottom right corner.
				if (noRight && noDown) || (noRight && !noDown && grid[curr.y+1][curr.x] != r) || (noDown && !noRight && grid[curr.y][curr.x+1] != r) || (!noRight && !noDown && grid[curr.y][curr.x+1] != r && grid[curr.y+1][curr.x] != r) || (!noRight && !noDown && grid[curr.y][curr.x+1] == r && grid[curr.y+1][curr.x] == r && grid[curr.y+1][curr.x+1] != r) {
					sides++
				}
			}

			score += area * sides
		}
	}

	println(score)
}
