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

func (p Point) Key() string {
	return string(p.x) + "," + string(p.y)
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
	antinodes := make(map[string]bool)

	for y, row := range grid {
		for x, r := range row {
			if r != '.' {
				p := Point{x, y}
				signals[r] = append(signals[r], p)
				antinodes[p.Key()] = true
			}
		}
	}

	for _, points := range signals {
		for i, p1 := range points {

			for j, p2 := range points {
				if i == j {
					continue
				}

				d := Distance(p1, p2)

				i := 1
				for {
					nodeX := p2.x + (d.x * i)
					nodeY := p2.y + (d.y * i)

					if nodeX < 0 || nodeX >= len(grid[0]) || nodeY < 0 || nodeY >= len(grid) {
						break
					}

					key := string(nodeX) + "," + string(nodeY)
					antinodes[key] = true
					i++
				}
			}
		}
	}

	fmt.Println(len(antinodes))
}
