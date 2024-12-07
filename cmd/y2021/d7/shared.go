package day7

import (
	"strconv"
	"strings"
)

func parseInput(lines []string) []int {
	var submarines []int
	for _, s := range strings.Split(lines[0], ",") {
		x, _ := strconv.Atoi(s)
		submarines = append(submarines, x)
	}
	return submarines
}
