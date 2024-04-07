package day9

import (
	"strconv"
	"strings"
)

func parseInput(lines []string) [][]int {
	var heat_map [][]int

	for _, line := range lines {
		var row []int
		for _, element := range strings.Split(line, "") {
			i, _ := strconv.Atoi(element)
			row = append(row, i)
		}
		heat_map = append(heat_map, row)
	}

	return heat_map

}

func localMin(heat_map [][]int, x int, y int) bool {
	if x != 0 && heat_map[y][x] >= heat_map[y][x-1] {
		return false
	} else if y != 0 && heat_map[y][x] >= heat_map[y-1][x] {
		return false
	} else if x != len(heat_map[0])-1 && heat_map[y][x] >= heat_map[y][x+1] {
		return false
	} else if y != len(heat_map)-1 && heat_map[y][x] >= heat_map[y+1][x] {
		return false
	}
	return true
}
