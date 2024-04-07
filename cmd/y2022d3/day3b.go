package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Set = map[string]bool

func line_to_set(line string) Set {
	m := make(map[string]bool)
	for _, c := range line {
		m[string(c)] = true
	}
	return m
}

func intersection(s1 Set, s2 Set) Set {
	s3 := make(Set)
	for k, _ := range s1 {
		if _, found := s2[k]; found {
			s3[k] = true
		}
	}
	return s3
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

	var score int
	var badges []string

	for i := 0; i < len(lines); i += 3 {
		s1 := line_to_set(lines[i])
		s2 := line_to_set(lines[i+1])
		s3 := line_to_set(lines[i+2])

		s4 := intersection(s1, intersection(s2, s3))

		for k, _ := range s4 {
			badges = append(badges, k)
		}
	}

	fmt.Println(badges)

	for _, s := range badges {
		b := s[0]
		if strings.ToLower(s) == s {
			score += int(b) - 96
		} else {
			score += int(b) - 38
		}
	}

	fmt.Println(score)
}
