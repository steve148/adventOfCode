package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(f)
	var grid [][]string
	for scan.Scan() {
		line := scan.Text()
		split := strings.Split(line, "")
		grid = append(grid, split)
	}

	accessibleRolls := 0
	for j := 0; j < len(grid); j++ {
		row := grid[j]
		for i := 0; i < len(row); i++ {
			if grid[j][i] != "@" {
				continue
			}

			count := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}

					j2 := j + dy
					i2 := i + dx

					if j2 < 0 || j2 >= len(grid) || i2 < 0 || i2 >= len(row) {
						continue
					}

					if grid[j2][i2] == "@" {
						count++
					}
				}
			}

			if count < 4 {
				accessibleRolls++
			}

		}
	}
	fmt.Println(accessibleRolls)

}
