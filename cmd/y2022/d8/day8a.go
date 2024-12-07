package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type key struct {
	row    int
	column int
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var trees [][]int

	for _, line := range lines {
		var row []int
		for _, s := range strings.Split(line, "") {
			i, _ := strconv.Atoi(s)
			row = append(row, i)
		}
		trees = append(trees, row)
	}

	width := len(trees[0])
	height := len(trees)

	visible_coordinates := make(map[key]bool)

	// Check left to right
	for i := 0; i < height; i++ {
		current_max := -1

		for j := 0; j < width; j++ {
			tree := trees[i][j]

			if tree > current_max {
				current_max = tree
				visible_coordinates[key{row: i, column: j}] = true
			}
		}
	}

	// Check right to left
	for i := 0; i < height; i++ {
		current_max := -1

		for j := width - 1; j >= 0; j-- {
			tree := trees[i][j]

			if tree > current_max {
				current_max = tree
				visible_coordinates[key{row: i, column: j}] = true
			}
		}
	}

	// Check top to bottom
	for i := 0; i < width; i++ {
		current_max := -1

		for j := 0; j < height; j++ {
			tree := trees[j][i]

			if tree > current_max {
				current_max = tree
				visible_coordinates[key{row: j, column: i}] = true
			}
		}
	}

	// Check bottom to top
	for i := 0; i < width; i++ {
		current_max := -1

		for j := height - 1; j >= 0; j-- {
			tree := trees[j][i]

			if tree > current_max {
				current_max = tree
				visible_coordinates[key{row: j, column: i}] = true
			}
		}
	}

	fmt.Println(len(visible_coordinates))
}
