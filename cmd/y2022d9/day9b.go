package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
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

	var knots []*coordinate
	num_knots := 10

	for i := 0; i < num_knots; i++ {
		knots = append(knots, &coordinate{x: 0, y: 0})
	}

	head := knots[0]
	tail := knots[num_knots-1]

	tail_visited := make(map[coordinate]bool)
	tail_visited[coordinate{x: tail.x, y: tail.y}] = true

	for _, line := range lines {
		split := strings.Split(line, " ")
		direction := split[0]
		steps, _ := strconv.Atoi(split[1])

		for i := 0; i < steps; i++ {
			// Move head
			switch direction {
			case "R":
				head.x++
			case "L":
				head.x--
			case "U":
				head.y--
			case "D":
				head.y++
			}

			// Move each knot in the chain
			for i := 1; i < num_knots; i++ {
				h := knots[i-1]
				t := knots[i]

				if (h.x == t.x-1 && h.y == t.y-2) || (h.x == t.x-2 && h.y == t.y-1) || (h.x == t.x-2 && h.y == t.y-2) {
					t.x--
					t.y--
				} else if h.x == t.x && h.y == t.y-2 {
					t.y--
				} else if (h.x == t.x+1 && h.y == t.y-2) || (h.x == t.x+2 && h.y == t.y-1) || (h.x == t.x+2 && h.y == t.y-2) {
					t.x++
					t.y--
				} else if h.x == t.x+2 && h.y == t.y {
					t.x++
				} else if (h.x == t.x+2 && h.y == t.y+1) || (h.x == t.x+1 && h.y == t.y+2) || (h.x == t.x+2 && h.y == t.y+2) {
					t.x++
					t.y++
				} else if h.x == t.x && h.y == t.y+2 {
					t.y++
				} else if (h.x == t.x-1 && h.y == t.y+2) || (h.x == t.x-2 && h.y == t.y+1) || (h.x == t.x-2 && h.y == t.y+2) {
					t.x--
					t.y++
				} else if h.x == t.x-2 && h.y == t.y {
					t.x--
				}
			}

			tail_visited[coordinate{x: tail.x, y: tail.y}] = true
		}
	}

	fmt.Println(len(tail_visited))
}
