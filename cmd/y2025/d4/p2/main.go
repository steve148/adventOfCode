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

	rollCoordinates := make(map[[2]int]bool)
	for j := 0; j < len(grid); j++ {
		for i := 0; i < len(grid[j]); i++ {
			if grid[j][i] == "@" {
				key := [2]int{j, i}
				rollCoordinates[key] = true
			}
		}
	}

	allRollsRemoved := false
	removedCount := 0
	for !allRollsRemoved {
		allRollsRemoved = true
		for key := range rollCoordinates {
			j, i := key[0], key[1]

			adjacentRolls := 0
			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if dx == 0 && dy == 0 {
						continue
					}

					j2, i2 := j+dy, i+dx
					key2 := [2]int{j2, i2}

					if rollCoordinates[key2] {
						adjacentRolls++
					}
				}
			}

			if adjacentRolls < 4 {
				removedCount++
				delete(rollCoordinates, key)
				allRollsRemoved = false
			}
		}
	}

	fmt.Println(removedCount)

}
