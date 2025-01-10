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
				switch r {
				case '#':
					row = append(row, '#')
					row = append(row, '#')
				case 'O':
					row = append(row, '[')
					row = append(row, ']')
				case '.':
					row = append(row, '.')
					row = append(row, '.')
				case '@':
					row = append(row, '@')
					row = append(row, '.')
				default:
					panic("Bad")
				}
			}
			grid = append(grid, row)
		} else {
			directionsSet = append(directionsSet, line)
		}
	}

	fmt.Println("Initial state")
	PrintGrid(grid)
	fmt.Println()

	for _, directions := range directionsSet {
		for _, dir := range directions {
			x, y := FindRobot(grid)
			ok := Check(grid, x, y, dir)

			if ok {
				grid = Swap(grid, x, y, dir)
			}

			// fmt.Printf("Move %c \n", dir)
			// PrintGrid(grid)
			// fmt.Println()
		}
	}

	totalScore := 0
	for y1, row := range grid {
		for x1, cell := range row {
			if cell != '[' {
				continue
			}
			score := GpsScore(x1, y1)
			totalScore += score
		}
	}
	fmt.Println("Total score:", totalScore)
}

func Check(grid [][]rune, x1, y1 int, dir rune) bool {
	dx, dy := InstructionToVector(dir)
	x2, y2 := x1+dx, y1+dy
	next := grid[y2][x2]

	if next == '#' {
		return false
	}
	if next == '.' {
		return true
	}

	if dx != 0 && (next == '[' || next == ']') {
		return Check(grid, x2, y2, dir)

	}
	if dy != 0 && next == '[' {
		left := Check(grid, x2, y2, dir)
		right := Check(grid, x2+1, y2, dir)
		return left && right
	}
	if dy != 0 && next == ']' {
		left := Check(grid, x2-1, y2, dir)
		right := Check(grid, x2, y2, dir)
		return left && right
	}

	panic("Bad")
}

func Swap(grid [][]rune, x1, y1 int, dir rune) [][]rune {
	dx, dy := InstructionToVector(dir)
	x2, y2 := x1+dx, y1+dy
	next := grid[y2][x2]

	if next == '#' {
		panic("No swap wall")
	}
	if next == '.' {
		SwapCell(grid, x1, y1, x2, y2)
		return grid
	}

	if dx != 0 && (next == '[' || next == ']') {
		grid := Swap(grid, x2, y2, dir)
		SwapCell(grid, x1, y1, x2, y2)
		return grid
	}

	if dy != 0 && next == '[' {
		grid1 := Swap(grid, x2, y2, dir)
		grid2 := Swap(grid1, x2+1, y2, dir)
		SwapCell(grid, x1, y1, x2, y2)
		return grid2
	}
	if dy != 0 && next == ']' {
		grid1 := Swap(grid, x2-1, y2, dir)
		grid2 := Swap(grid1, x2, y2, dir)
		SwapCell(grid, x1, y1, x2, y2)
		return grid2
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

func SwapCell(grid [][]rune, x1, y1, x2, y2 int) {
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
