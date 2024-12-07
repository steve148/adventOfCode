package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Part1() {
	file, err := os.Open("./day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	unique_codes := 0
	for _, line := range lines {
		split := strings.Fields(line)
		output := split[11:]

		for _, code := range output {
			if len(code) == 2 || len(code) == 3 || len(code) == 4 || len(code) == 7 {
				unique_codes += 1
			}
		}
	}

	fmt.Println(unique_codes)
}
