package day14

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Part1() {
	file, err := os.Open("./day14/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	pairs, rules, counts := parseInput(lines)

	for step := 0; step < 10; step++ {
		newPairs := make(map[string]int)
		for pair, count := range pairs {
			newPairs[pair[0:1]+rules[pair]] += count
			newPairs[rules[pair]+pair[1:2]] += count
			counts[rules[pair]] += count
		}
		pairs = newPairs
	}

	min := int(^uint(0) >> 1)
	max := 0
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	fmt.Println(max - min)

}
