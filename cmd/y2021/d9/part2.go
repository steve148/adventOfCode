package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func findNeighbours(heat_map [][]int, x int, y int) [][]int {
	var neighbours [][]int
	if x != 0 && heat_map[y][x-1] != 9 {
		neighbours = append(neighbours, []int{x - 1, y})
	}
	if y != 0 && heat_map[y-1][x] != 9 {
		neighbours = append(neighbours, []int{x, y - 1})
	}
	if x != len(heat_map[0])-1 && heat_map[y][x+1] != 9 {
		neighbours = append(neighbours, []int{x + 1, y})
	}
	if y != len(heat_map)-1 && heat_map[y+1][x] != 9 {
		neighbours = append(neighbours, []int{x, y + 1})
	}
	return neighbours
}

func basinSize(heat_map [][]int, x int, y int) int {
	count := 0

	var seen_locations [][]bool
	for j := 0; j < len(heat_map); j++ {
		var row []bool
		for i := 0; i < len(heat_map[0]); i++ {
			row = append(row, false)
		}
		seen_locations = append(seen_locations, row)
	}

	var queue [][]int
	queue = append(queue, findNeighbours(heat_map, x, y)...)

	for len(queue) > 0 {
		i := queue[0][0]
		j := queue[0][1]
		queue = queue[1:]

		if !seen_locations[j][i] {
			count += 1
			seen_locations[j][i] = true

			queue = append(queue, findNeighbours(heat_map, i, j)...)
		}
	}

	return count
}

func Part2() {
	file, err := os.Open("./day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	heatMap := parseInput(lines)

	var basinSizes []int
	for y := 0; y < len(heatMap); y++ {
		for x := 0; x < len(heatMap[y]); x++ {
			if localMin(heatMap, x, y) {
				size := basinSize(heatMap, x, y)
				basinSizes = append(basinSizes, size)
			}
		}
	}

	sort.Ints(basinSizes)
	largestBasins := basinSizes[len(basinSizes)-3:]
	fmt.Println(largestBasins[0] * largestBasins[1] * largestBasins[2])
}
