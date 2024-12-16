package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func Distance(p1, p2 Point) Point {
	return Point{p2.x - p1.x, p2.y - p1.y}
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
		row := []rune(line)
		grid = append(grid, row)
	}

	signals := make(map[rune][]Point)

	for y, row := range grid {
		for x, r := range row {
			if r != '.' {
				signals[r] = append(signals[r], Point{x, y})
			}
		}
	}

	antinodes := make(map[string]bool)
	for _, points := range signals {
		for i, p1 := range points {

			for j, p2 := range points {
				if i == j {
					continue
				}

				d := Distance(p1, p2)
				nodeX := p2.x + d.x
				nodeY := p2.y + d.y

				if nodeX < 0 || nodeX >= len(grid[0]) || nodeY < 0 || nodeY >= len(grid) {
					continue
				}

				key := string(nodeX) + "," + string(nodeY)
				antinodes[key] = true
			}
		}
	}

	fmt.Println(len(antinodes))
}
