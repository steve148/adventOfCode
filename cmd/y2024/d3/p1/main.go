package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	b, err := os.ReadFile("./data.txt")
	if err != nil {
		panic(err)
	}

	s := string(b)

	fmt.Println(s)

	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllString(s, -1)

	sum := 0
	for _, m := range matches {
		for _, submatch := range r.FindAllStringSubmatch(m, -1) {
			n, err := strconv.Atoi(submatch[1])
			if err != nil {
				panic(err)
			}

			m, err := strconv.Atoi(submatch[2])
			if err != nil {
				panic(err)
			}

			sum += n * m
		}
	}
	fmt.Println(sum)

}
