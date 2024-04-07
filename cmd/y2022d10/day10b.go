package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var commands []int
	for _, line := range lines {
		commands = append(commands, 0)
		if strings.HasPrefix(line, "addx") {
			split := strings.Split(line, " ")
			x, _ := strconv.Atoi(split[1])
			commands = append(commands, x)
		}
	}

	var screen []string
	register := 1
	for i, add := range commands {
		column := i % 40

		if register == column || register == column-1 || register == column+1 {
			screen = append(screen, "#")
		} else {
			screen = append(screen, ".")
		}

		register += add
	}

	for i, s := range screen {
		if i%40 == 0 {
			fmt.Println()
		}
		fmt.Print(s)
	}
	fmt.Println()
}
