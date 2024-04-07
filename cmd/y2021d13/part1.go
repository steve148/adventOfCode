package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1() {
	file, err := os.Open("./day13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	coordinates, instructions, maxX, maxY := parseInput(lines)

	grid := initGrid(maxX+1, maxY+1, coordinates)

	for _, instruction := range instructions[:1] {
		if instruction.dimension == 'y' {
			grid = foldUp(grid, instruction.position)
		} else {
			grid = foldLeft(grid, instruction.position)
		}

	}

	count := 0
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[0]); i++ {
			if grid[j][i] == "#" {
				count++

			}
		}
	}

	fmt.Println(count)
}
