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

	var duplicates []string

	for _, line := range lines {
		m := make(map[string]bool)
		for i := 0; i < len(line)/2; i++ {
			m[string(line[i])] = true
		}
		for i := len(line) / 2; i < len(line); i++ {
			c := string(line[i])
			if _, ok := m[c]; ok {
				duplicates = append(duplicates, c)
				break
			}
		}
	}

	score := 0
	fmt.Println("a", 'a') // 97
	fmt.Println("A", 'A') // 65

	for _, s := range duplicates {
		b := s[0]
		if strings.ToLower(s) == s {
			score += int(b) - 96
		} else {
			score += int(b) - 38
		}
	}
	fmt.Println(score)
}
