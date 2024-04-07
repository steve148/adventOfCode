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

	strength := 0
	register := 1
	for i, add := range commands {
		if (i+1-20)%40 == 0 {
			fmt.Println(i+1, register)
			strength += (i + 1) * register
		}

		register += add
	}

	fmt.Println(strength)
}
