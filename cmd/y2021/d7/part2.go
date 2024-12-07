package day7

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
)

func Part2() {
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

	sum := 0.0
	for _, position := range positions {
		sum += float64(position)
	}

	align_point := int(sum / float64(len(positions)))
	fmt.Println("Align point", align_point)

	fuel := 0
	for _, position := range positions {
		difference := int(math.Abs(float64(position - align_point)))
		fuel += (difference * (difference + 1)) / 2
		if (difference*(difference+1))/2 != factorial(difference) {
			fmt.Println(difference, (difference*(difference+1))/2, factorial(difference))
		}
	}

	fmt.Println("Fuel", fuel)
}

func factorial(n int) int {
	if n == 0 {
		return 0
	}
	return n + factorial(n-1)
}
