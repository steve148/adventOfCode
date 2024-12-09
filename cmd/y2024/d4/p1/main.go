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

			searchUp := y-3 >= 0
			searchRight := x+3 <= len(row)-1
			searchDown := y+3 <= len(grid)-1
			searchLeft := x-3 >= 0

			fmt.Println(searchUp, searchRight, searchDown, searchLeft)

			if searchUp && grid[y][x] == 'X' && grid[y-1][x] == 'M' && grid[y-2][x] == 'A' && grid[y-3][x] == 'S' {
				xmasCount++
			}
			if searchUp && searchRight && grid[y][x] == 'X' && grid[y-1][x+1] == 'M' && grid[y-2][x+2] == 'A' && grid[y-3][x+3] == 'S' {
				xmasCount++
			}
			if searchRight && grid[y][x] == 'X' && grid[y][x+1] == 'M' && grid[y][x+2] == 'A' && grid[y][x+3] == 'S' {
				xmasCount++
			}
			if searchRight && searchDown && grid[y][x] == 'X' && grid[y+1][x+1] == 'M' && grid[y+2][x+2] == 'A' && grid[y+3][x+3] == 'S' {
				xmasCount++
			}
			if searchDown && grid[y][x] == 'X' && grid[y+1][x] == 'M' && grid[y+2][x] == 'A' && grid[y+3][x] == 'S' {
				xmasCount++
			}
			if searchLeft && searchDown && grid[y][x] == 'X' && grid[y+1][x-1] == 'M' && grid[y+2][x-2] == 'A' && grid[y+3][x-3] == 'S' {
				xmasCount++
			}
			if searchLeft && grid[y][x] == 'X' && grid[y][x-1] == 'M' && grid[y][x-2] == 'A' && grid[y][x-3] == 'S' {
				xmasCount++
			}
			if searchLeft && searchUp && grid[y][x] == 'X' && grid[y-1][x-1] == 'M' && grid[y-2][x-2] == 'A' && grid[y-3][x-3] == 'S' {
				xmasCount++
			}

		}
	}

	fmt.Println(xmasCount)
}
