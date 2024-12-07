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

	min_x := 1000
	max_x := 0
	max_y := 0

	fmt.Println("Read input as pairs of coordinates")
	var rocks [][]*Coordinate
	for _, line := range lines {
		var prev *Coordinate

		for _, s := range strings.Split(line, " -> ") {
			split := strings.Split(s, ",")
			x, _ := strconv.Atoi(split[0])
			y, _ := strconv.Atoi(split[1])

			if x < min_x {
				min_x = x
			}
			if x > max_x {
				max_x = x
			}
			if y > max_y {
				max_y = y
			}

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

	width := max_x - min_x + 1
	height := max_y + 1
	fmt.Println(min_x, max_x, max_y, width, height)

	fmt.Println("Normalize coordinate values.")
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

	// print_grid(grid)

	var count int
	reached_void := false

	fmt.Println("Determine sand count.")
	for {
		p := &Coordinate{norm_x: 500 - min_x, norm_y: 0}

		for {
			if p.norm_x == 0 || p.norm_x == len(grid[0])-1 || p.norm_y == len(grid)-1 {
				reached_void = true
				break
			} else if grid[p.norm_y+1][p.norm_x] == "." {
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

		// print_grid(grid)

		if reached_void == true {
			break
		}
		count++
	}

	print_grid(grid)

	fmt.Println(count)
}
