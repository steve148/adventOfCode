package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1() {
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

	heat_map := parseInput(lines)

	risk := 0
	for y := 0; y < len(heat_map); y++ {
		for x := 0; x < len(heat_map[y]); x++ {
			if localMin(heat_map, x, y) {
				risk += heat_map[y][x] + 1
			}
		}
	}
	fmt.Println(risk)

}
