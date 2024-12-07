package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func Part1() {
	file, err := os.Open("./day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	positions := parseInput(lines)
	sort.Ints(positions)
	fmt.Println(positions)

	align_point := positions[len(positions)/2]

	fuel := 0
	for _, position := range positions {
		fuel += int(math.Abs(float64(position - align_point)))
	}

	fmt.Println(fuel)
}
