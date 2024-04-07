package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getCoordinates(vent_lines [][][]int) [][]int {
	var coordinates [][]int
	for _, vent := range vent_lines {
		x1 := vent[0][0]
		y1 := vent[0][1]
		x2 := vent[1][0]
		y2 := vent[1][1]

		if x1 == x2 {
			// horizontal
			ys := getRangeOfInts(y1, y2)
			for _, y := range ys {
				coordinates = append(coordinates, []int{x1, y})
			}
		} else if y1 == y2 {
			// vertical
			xs := getRangeOfInts(x1, x2)
			for _, x := range xs {
				coordinates = append(coordinates, []int{x, y1})
			}
		}
	}
	return coordinates
}

func Part1() {
	file, err := os.Open("./day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	vent_lines, max_x, max_y := parseInput(lines)

	grid := populateGrid(max_x, max_y)

	coordinates := getCoordinates(vent_lines)

	marked_grid := markGrid(coordinates, grid)
	printGrid(marked_grid)

	sum := sumGrid(marked_grid)

	fmt.Println(sum)
}
