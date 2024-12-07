package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y, norm_x, norm_y int
}

func print_grid(grid [][]string) {
	fmt.Println("Grid")
	for _, row := range grid {
		for _, s := range row {
			fmt.Print(s)
		}
		fmt.Println()
	}
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

	fmt.Println("Read input as pairs of coordinates")
	var rocks [][]*Coordinate
	for _, line := range lines {
		var prev *Coordinate

		for _, s := range strings.Split(line, " -> ") {
			split := strings.Split(s, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			c := &Coordinate{x: x, y: y}

			if prev != nil {
				if prev.x < c.x || prev.y < c.y {
					rocks = append(rocks, []*Coordinate{prev, c})
				} else {
					rocks = append(rocks, []*Coordinate{c, prev})
				}
			}

			prev = c
		}
	}

	fmt.Println("Find grid height.")
	max_y := 0
	for _, rock := range rocks {
		if rock[1].y > max_y {
			max_y = rock[1].y
		}
	}
	// Add two for floor.
	max_y += 2
	height := max_y + 1
	fmt.Println("Max", max_y, "Height", height)

	fmt.Println("Find grid width.")
	min_x := 1000
	max_x := 0

	// Project each x coordinate downward to determine min and max.
	start := &Coordinate{x: 500, y: 0}
	if start.x-(max_y-start.y) < min_x {
		min_x = start.x - (max_y - start.y)
	}
	if start.x+(max_y-start.y) > max_x {
		max_x = start.x + (max_y - start.y)
	}

	for _, rock := range rocks {
		if rock[0].x-(max_y-rock[0].y) < min_x {
			min_x = rock[0].x - (max_y - rock[0].y)
		}
		if rock[1].x+(max_y-rock[0].y) > max_x {
			max_x = rock[1].x + (max_y - rock[0].y)
		}
	}
	width := max_x - min_x + 1
	fmt.Println("Min", min_x, "Max", max_x, "Width", width)

	fmt.Println("Normalize coordinate values.")
	start.norm_x = start.x - min_x
	start.norm_y = start.y
	for _, rock := range rocks {
		for _, c := range rock {
			c.norm_x = c.x - min_x
			c.norm_y = c.y
		}
	}

	fmt.Println("Build 2D grid.")
	var grid [][]string
	for i := 0; i < height; i++ {
		var row []string
		for j := 0; j < width; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	fmt.Println("Populate grid with rocks")
	for _, rock := range rocks {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				left := rock[0].norm_x
				right := rock[1].norm_x
				top := rock[0].norm_y
				bottom := rock[1].norm_y
				if top <= i && i <= bottom && left <= j && j <= right {
					grid[i][j] = "#"
				}
			}
		}
	}

	fmt.Println("Bottom row is all rocks.")
	for i := 0; i < width; i++ {
		grid[height-1][i] = "#"
	}

	var count int
	reached_void := false

	fmt.Println("Determine sand count.")
	for {
		p := &Coordinate{norm_x: start.norm_x, norm_y: start.norm_y}

		if grid[p.norm_y][p.norm_x] == "o" {
			break
		}

		for {
			if grid[p.norm_y+1][p.norm_x] == "." {
				p.norm_y++
			} else if grid[p.norm_y+1][p.norm_x-1] == "." {
				p.norm_y++
				p.norm_x--
			} else if grid[p.norm_y+1][p.norm_x+1] == "." {
				p.norm_y++
				p.norm_x++
			} else {
				grid[p.norm_y][p.norm_x] = "o"
				break
			}
		}

		if reached_void == true {
			break
		}
		count++
	}

	fmt.Println(count)
}
