package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var elf_rations [][]int
	var current_elf_rations []int
	for _, line := range lines[1:] {
		if line == "" {
			elf_rations = append(elf_rations, current_elf_rations)
			current_elf_rations = nil
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			current_elf_rations = append(current_elf_rations, i)
		}
	}
	elf_rations = append(elf_rations, current_elf_rations)

	var elf_ration_sum []int
	for _, elf_ration := range elf_rations {
		ration_sum := 0
		for _, ration := range elf_ration {
			ration_sum += ration
		}
		elf_ration_sum = append(elf_ration_sum, ration_sum)
	}

	max := 0
	for _, sum := range elf_ration_sum {
		if max < sum {
			max = sum
		}
	}
	fmt.Println(max)
}
