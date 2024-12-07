package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fptr := flag.String("fpath", "data.txt", "file path to read from")
	flag.Parse()

	file, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	line_scanner := bufio.NewScanner(file)

	line_scanner.Split(bufio.ScanLines)

	var lines []string
	for line_scanner.Scan() {
		lines = append(lines, line_scanner.Text())
	}

	var result int

	for _, line := range lines {
		fmt.Println(line)

		for i, c := range line {
			fmt.Println(line[i], c)

		}

	}
	fmt.Println("Year 2023 Day 1 Part 1 Result: ", result)

}
