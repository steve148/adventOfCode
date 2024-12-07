package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func countFlashes(octopi [][]int) int {
	count := 0
	height := len(octopi)
	width := len(octopi[0])

	for k := 1; k <= 100; k++ {
		// Increment all elements
		for j := 0; j < height; j++ {
			for i := 0; i < width; i++ {
				octopi[j][i]++
			}
		}

		anyFlashes := true
		for anyFlashes {
			anyFlashes = false

			for j := 0; j < height; j++ {
				for i := 0; i < width; i++ {
					if octopi[j][i] > 9 {
						count++
						anyFlashes = true
						octopi[j][i] = 0

						if i != 0 && j != 0 && octopi[j-1][i-1] != 0 { // up left
							octopi[j-1][i-1]++
						}
						if j != 0 && octopi[j-1][i] != 0 { // up
							octopi[j-1][i]++
						}
						if i != width-1 && j != 0 && octopi[j-1][i+1] != 0 { // up right
							octopi[j-1][i+1]++
						}
						if i != 0 && octopi[j][i-1] != 0 { // left
							octopi[j][i-1]++
						}
						if i != width-1 && octopi[j][i+1] != 0 { // right
							octopi[j][i+1]++
						}
						if i != 0 && j != height-1 && octopi[j+1][i-1] != 0 { // bottom left
							octopi[j+1][i-1]++
						}
						if j != height-1 && octopi[j+1][i] != 0 { // bottom
							octopi[j+1][i]++
						}
						if i != width-1 && j != height-1 && octopi[j+1][i+1] != 0 { // bottom right
							octopi[j+1][i+1]++
						}
					}
				}
			}
		}
		fmt.Println("Step", k, count)
		for _, row := range octopi {
			fmt.Println(row)
		}
	}

	return count
}

func Part1() {
	file, err := os.Open("./day11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var octopi [][]int
	for _, line := range lines {
		var row []int
		for _, element := range strings.Split(line, "") {
			i, _ := strconv.Atoi(element)
			row = append(row, i)
		}
		octopi = append(octopi, row)
	}

	flashes := countFlashes(octopi)

	fmt.Println(flashes)
}
