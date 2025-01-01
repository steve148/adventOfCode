package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	var grid [][]rune
	var directionsSet []string
	part := 0
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()

		if line == "" {
			part++
		} else if part == 0 {
			var row []rune
			for _, r := range line {
				row = append(row, r)
			}
			grid = append(grid, row)
		} else {
			directionsSet = append(directionsSet, line)
		}
	}

	// fmt.Println("Initial state")
	// PrintGrid(grid)
	// fmt.Println()

	for _, directions := range directionsSet {
		for _, dir := range directions {
			x, y := FindRobot(grid)
			_, grid = CheckAndSwap(grid, x, y, dir)

			// fmt.Printf("Move %c \n", dir)
			// PrintGrid(grid)
			// fmt.Println()
		}
	}

	totalScore := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'O' {
				totalScore += GpsScore(x, y)
			}
		}
	}
	fmt.Println("Total score:", totalScore)
}

func CheckAndSwap(grid [][]rune, x1, y1 int, dir rune) (bool, [][]rune) {
	dx, dy := InstructionToVector(dir)
	x2, y2 := x1+dx, y1+dy
	if grid[y2][x2] == '#' {
		return false, grid
	}
	if grid[y2][x2] == '.' {
		Swap(grid, x1, y1, x2, y2)
		return true, grid
	}
	if grid[y2][x2] == 'O' {
		check, grid := CheckAndSwap(grid, x2, y2, dir)
		if check {
			Swap(grid, x1, y1, x2, y2)
		}
		return check, grid
	}
	panic("Bad")
}

func FindRobot(grid [][]rune) (int, int) {
	for j, row := range grid {
		for i, cell := range row {
			if cell == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func InstructionToVector(instruction rune) (int, int) {
	switch instruction {
	case '<':
		return -1, 0
	case '>':
		return 1, 0
	case '^':
		return 0, -1
	case 'v':
		return 0, 1
	default:
		return 0, 0
	}
}

func Swap(grid [][]rune, x1, y1, x2, y2 int) {
	grid[y1][x1], grid[y2][x2] = grid[y2][x2], grid[y1][x1]
}

func PrintGrid(grid [][]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Print(string(cell))
		}
		fmt.Println()
	}
}

func GpsScore(x, y int) int {
	return y*100 + x
}
