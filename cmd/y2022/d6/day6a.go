package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	line := lines[0]
	num_distinct_chars := 4
	for i := num_distinct_chars; i < len(line)-1; i++ {
		four_characters := line[i-num_distinct_chars : i]
		set := make(map[rune]bool)
		has_duplicate := false

		for _, r := range four_characters {
			_, exists := set[r]
			if exists {
				has_duplicate = true
			}
			set[r] = true
		}

		if !has_duplicate {
			fmt.Println(i)
			break
		}
	}
}
