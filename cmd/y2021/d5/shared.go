package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func arrToi(a []string) []int {
	var nums []int
	for _, s := range a {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}
	return nums
}

func parseInput(lines []string) ([][][]int, int, int) {
	max_x := 0
	max_y := 0
	var vent_lines [][][]int

	for _, line := range lines {
		split := strings.Split(line, "->")
		var coordinates [][]int
		for _, s := range split {
			coordinate := arrToi(strings.Split(strings.TrimSpace(s), ","))

			if coordinate[0] > max_x {
				max_x = coordinate[0]
			}

			if coordinate[1] > max_y {
				max_y = coordinate[1]
			}

			coordinates = append(coordinates, coordinate)
		}
		vent_lines = append(vent_lines, coordinates)
	}
	return vent_lines, max_x, max_y
}

func getRangeOfInts(n1 int, n2 int) []int {
	var min, max int
	if n1 > n2 {
		max = n1
		min = n2
	} else {
		max = n2
		min = n1
	}

	var n []int
	for i := min; i <= max; i++ {
		n = append(n, i)
	}
	return n
}

func populateGrid(x int, y int) [][]int {
	var grid [][]int
	for i := 0; i < y+1; i++ {
		var row []int
		for j := 0; j < x+1; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}
	return grid
}

func markGrid(coordinates [][]int, grid [][]int) [][]int {
	for _, coordinate := range coordinates {
		x := coordinate[0]
		y := coordinate[1]
		grid[y][x] += 1

	}
	return grid
}

func sumGrid(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				count += 1
			}
		}
	}
	return count
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}
