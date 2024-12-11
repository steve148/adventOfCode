package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(readFile)
	scan.Split(bufio.ScanLines)

	var grid [][]rune
	var seen [][]bool

	y := -1
	incrX := 0
	incrY := -1
	var guardX, guardY int
	for scan.Scan() {
		line := scan.Text()
		y++

		var gridRow []rune
		var seenRow []bool

		for x, r := range line {
			if r == '^' {
				guardX = x
				guardY = y
				seenRow = append(seenRow, true)
			} else {
				seenRow = append(seenRow, false)
			}
			gridRow = append(gridRow, r)
		}
		grid = append(grid, gridRow)
		seen = append(seen, seenRow)
	}

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
			seen[guardY][guardX] = true
		}

	}

	seenCount := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if seen[y][x] {
				seenCount++
			}
		}
	}

	fmt.Println(seenCount)
}
