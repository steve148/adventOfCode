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
	x int
	y int
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

	// Pick a large size for grid.
	tail_coordinates := make(map[key]bool)

	hx, hy, tx, ty := 0, 0, 0, 0
	tail_coordinates[key{x: tx, y: ty}] = true

	for _, line := range lines {
		split := strings.Split(line, " ")
		direction := split[0]
		steps, _ := strconv.Atoi(split[1])

		for i := 0; i < steps; i++ {

			// Move head
			switch direction {
			case "R":
				hx++
			case "L":
				hx--
			case "U":
				hy--
			case "D":
				hy++
			}

			if (hx == tx-1 && hy == ty-2) || (hx == tx-2 && hy == ty-1) {
				tx--
				ty--
				tail_coordinates[key{x: tx, y: ty}] = true
			} else if hx == tx && hy == ty-2 {
				ty--
				tail_coordinates[key{x: tx, y: ty}] = true
			} else if (hx == tx+1 && hy == ty-2) || (hx == tx+2 && hy == ty-1) {
				tx++
				ty--
				tail_coordinates[key{x: tx, y: ty}] = true
			} else if hx == tx+2 && hy == ty {
				tx++
				tail_coordinates[key{x: tx, y: ty}] = true
			} else if (hx == tx+2 && hy == ty+1) || (hx == tx+1 && hy == ty+2) {
				tx++
				ty++
				tail_coordinates[key{x: tx, y: ty}] = true
			} else if hx == tx && hy == ty+2 {
				ty++
				tail_coordinates[key{x: tx, y: ty}] = true
			} else if (hx == tx-1 && hy == ty+2) || (hx == tx-2 && hy == ty+1) {
				tx--
				ty++
				tail_coordinates[key{x: tx, y: ty}] = true
			} else if hx == tx-2 && hy == ty {
				tx--
				tail_coordinates[key{x: tx, y: ty}] = true
			}
		}
	}

	fmt.Println(len(tail_coordinates))
}
