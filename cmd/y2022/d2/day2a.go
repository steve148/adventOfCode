package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

	score := 0
	for _, line := range lines {
		split := strings.Fields(line)
		opponent := split[0]

		var player string
		switch split[1] {
		case "X":
			player = "A"
			score += 1
		case "Y":
			player = "B"
			score += 2
		case "Z":
			player = "C"
			score += 3
		}

		if player == opponent {
			score += 3
		}
		if (opponent == "A" && player == "B") || (opponent == "B" && player == "C") || (opponent == "C" && player == "A") {
			score += 6
		}
	}
	fmt.Println(score)
}
