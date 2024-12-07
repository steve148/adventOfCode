package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
		first_num := -1
		second_num := -1

		for _, c := range line {
			i, err := strconv.Atoi(string(c))
			if err == nil {
				if first_num == -1 {
					first_num = i
				} else {
					second_num = i
				}
			}
		}

		if second_num == -1 {
			second_num = first_num
		}

		result = result + first_num*10 + second_num
	}
	fmt.Println("Year 2023 Day 1 Part 1 Result: ", result)

}
