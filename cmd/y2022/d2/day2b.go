package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func score_opponent_result(opponent string, result string) int {
	if opponent == "A" {
		if result == "X" {
			return 3
		} else if result == "Y" {
			return 1
		} else {
			return 2
		}
	} else if opponent == "B" {
		if result == "X" {
			return 1
		} else if result == "Y" {
			return 2
		} else {
			return 3
		}
	} else {
		if result == "X" {
			return 2
		} else if result == "Y" {
			return 3
		} else {
			return 1
		}
	}
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

	score := 0
	for _, line := range lines {
		split := strings.Fields(line)
		opponent := split[0]
		result := split[1]

		if result == "Y" {
			score += 3
		} else if result == "Z" {
			score += 6
		}

		score += score_opponent_result(opponent, result)
	}
	fmt.Println(score)
}
