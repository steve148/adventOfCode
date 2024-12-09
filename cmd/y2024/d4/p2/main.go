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

	for scan.Scan() {
		line := scan.Text()
		var row []rune
		for _, r := range line {
			row = append(row, r)
		}
		grid = append(grid, row)
	}

	xmasCount := 0
	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'A' && x > 0 && x < len(row)-1 && y > 0 && y < len(grid)-1 {
				if grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S' && grid[y+1][x-1] == 'M' && grid[y-1][x+1] == 'S' {
					xmasCount++
				}
				if grid[y-1][x-1] == 'M' && grid[y+1][x+1] == 'S' && grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S' {
					xmasCount++
				}
				if grid[y+1][x+1] == 'M' && grid[y-1][x-1] == 'S' && grid[y-1][x+1] == 'M' && grid[y+1][x-1] == 'S' {
					xmasCount++
				}
				if grid[y+1][x+1] == 'M' && grid[y-1][x-1] == 'S' && grid[y+1][x-1] == 'M' && grid[y-1][x+1] == 'S' {
					xmasCount++
				}
			}
		}
	}

	fmt.Println(xmasCount)
}
