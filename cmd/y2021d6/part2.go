package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part2() {
	file, err := os.Open("./day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	state := parseInput(lines)
	printState(state, 0)

	for day := 1; day <= 256; day++ {
		state = updateState(state)
		printState(state, day)
	}

	fmt.Println(sumState(state))
}
