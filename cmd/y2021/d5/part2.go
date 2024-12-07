package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getAllCoordinates(vent_lines [][][]int) [][]int {
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
		} else if x1 < x2 && y1 < y2 {
			// right and up
			for i := 0; i <= x2-x1; i++ {
				coordinates = append(coordinates, []int{x1 + i, y1 + i})
			}
		} else if x1 < x2 && y1 > y2 {
			// right and down
			for i := 0; i <= x2-x1; i++ {
				coordinates = append(coordinates, []int{x1 + i, y1 - i})
			}
		} else if x1 > x2 && y1 < y2 {
			// left and up
			for i := 0; i <= x1-x2; i++ {
				coordinates = append(coordinates, []int{x1 - i, y1 + i})
			}

		} else if x1 > x2 && y1 > y2 {
			// left and down
			for i := 0; i <= x1-x2; i++ {
				coordinates = append(coordinates, []int{x1 - i, y1 - i})
			}
		}

	}
	return coordinates
}

func Part2() {
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

	coordinates := getAllCoordinates(vent_lines)

	marked_grid := markGrid(coordinates, grid)
	printGrid(marked_grid)

	sum := sumGrid(marked_grid)

	fmt.Println(sum)
}
