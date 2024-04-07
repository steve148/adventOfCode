package day13

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	dimension rune
	position  int
}

func parseInput(lines []string) (coordinates [][]int, instructions []instruction, maxX int, maxY int) {
	isCoordinate := true
	for _, line := range lines {
		if line == "" {
			isCoordinate = false
		} else if isCoordinate {
			split := strings.Split(line, ",")
			x, _ := strconv.Atoi(split[0])
			if x > maxX {
				maxX = x
			}

			y, _ := strconv.Atoi(split[1])
			if y > maxY {
				maxY = y
			}

			coordinates = append(coordinates, []int{x, y})
		} else {
			var dimension rune
			var position int
			_, err := fmt.Sscanf(line, "fold along %c=%d", &dimension, &position)
			if err != nil {
				log.Fatal(err)
			}
			instructions = append(instructions, instruction{dimension: dimension, position: position})
		}
	}

	return
}

func initGrid(width, height int, coordinates [][]int) (grid [][]string) {
	for j := 0; j < height; j++ {
		var row []string
		for i := 0; i < width; i++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for _, coordinate := range coordinates {
		x := coordinate[0]
		y := coordinate[1]
		grid[y][x] = "#"
	}
	return
}

func foldUp(grid [][]string, position int) (newGrid [][]string) {
	height := len(grid)

	for j := 0; j < position; j++ {
		var row []string
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == "#" || grid[height-1-j][i] == "#" {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}
		newGrid = append(newGrid, row)
	}
	return
}

func foldLeft(grid [][]string, position int) (newGrid [][]string) {
	width := len(grid[0])

	for j := 0; j < len(grid); j++ {
		var row []string
		for i := 0; i < position; i++ {
			if grid[j][i] == "#" || grid[j][width-1-i] == "#" {
				row = append(row, "#")
			} else {
				row = append(row, ".")
			}
		}
		newGrid = append(newGrid, row)
	}
	return
}
