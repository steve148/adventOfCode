package main

import (
	"bufio"
	"fmt"
	"os"
)

func GenSeen(grid [][]rune) [][]int {
	var seen [][]int
	for _, row := range grid {
		var seenRow []int
		for _, r := range row {
			if r == '^' {
				seenRow = append(seenRow, 1)
			} else {
				seenRow = append(seenRow, 0)
			}
		}
		seen = append(seen, seenRow)
	}
	return seen
}

func FindGuard(grid [][]rune) (int, int) {
	for y, row := range grid {
		for x, r := range row {
			if r == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}

func HasLoop(grid [][]rune) bool {
	seen := GenSeen(grid)
	guardX, guardY := FindGuard(grid)
	incrX := 0
	incrY := -1

	for {
		nextX := guardX + incrX
		nextY := guardY + incrY
		if nextX < 0 || nextX == len(grid[0]) || nextY < 0 || nextY == len(grid) {
			// Out of bounds, end of path.
			break
		} else if grid[nextY][nextX] == '#' {
			// Rotate right 90 degrees
			incrX, incrY = -incrY, incrX
		} else if seen[nextY][nextX] == 4 {
			return true
		} else {
			// Move forward one space
			guardX = nextX
			guardY = nextY
			seen[guardY][guardX]++
		}
	}
	return false
}

func main() {
	readFile, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(readFile)
	scan.Split(bufio.ScanLines)

	var grid [][]rune
	for scan.Scan() {
		line := scan.Text()
		var gridRow []rune

		for _, r := range line {
			gridRow = append(gridRow, r)
		}
		grid = append(grid, gridRow)
	}

	seen := GenSeen(grid)
	guardX, guardY := FindGuard(grid)
	incrX := 0
	incrY := -1

	// Find path that guard would take if there were no extra blockers.
	for {
		nextX := guardX + incrX
		nextY := guardY + incrY
		if nextX < 0 || nextX == len(grid[0]) || nextY < 0 || nextY == len(grid) {
			// Out of bounds, end of path.
			break
		} else if grid[nextY][nextX] == '#' {
			// Rotate right 90 degrees
			incrX, incrY = -incrY, incrX
		} else {
			// Move forward one space
			guardX = nextX
			guardY = nextY
			seen[guardY][guardX]++
		}
	}

	// Brute force through all possible paths.
	loopCount := 0
	for y := range grid {
		for x := range grid[y] {
			if seen[y][x] == 0 {
				// Guard won't see this spot, can't be a loop.
				continue
			}

			if grid[y][x] == '^' {
				// Guard's starting position, can't be a loop.
				continue
			}

			// Set a blocker in the spot.
			grid[y][x] = '#'

			if HasLoop(grid) {
				loopCount++
			}

			grid[y][x] = '.'
		}
	}
	fmt.Println(loopCount)
}
