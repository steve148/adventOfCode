package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	arg := os.Args[1]
	iter, err := strconv.Atoi(arg)
	if err != nil {
		panic(err)
	}

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var line string
	if scanner.Scan() {
		line = scanner.Text()
	}

	int_regex := regexp.MustCompile(`\d+`)
	matches := int_regex.FindAllString(line, -1)

	var stones []int
	for _, s := range matches {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		stones = append(stones, i)
	}

	stoneMap := make(map[int]int)
	for _, stone := range stones {
		stoneMap[stone]++
	}

	for i := 0; i < iter; i++ {
		stoneMap = Blink(stoneMap)
	}

	fmt.Println(NumStones(stoneMap))
}

func NumStones(stoneMap map[int]int) int {
	var total int
	for _, count := range stoneMap {
		total += count
	}
	return total
}

func Blink(stones map[int]int) map[int]int {
	result := make(map[int]int)
	for stone, count := range stones {
		if stone == 0 {
			// Rule 1: Swap 0 for 1
			result[1] += count
			continue
		}

		str := strconv.Itoa(stone)
		if len(str)%2 == 0 {
			// Rule 2: Split into two stones
			half := len(str) / 2
			left, err := strconv.Atoi(str[:half])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(str[half:])
			if err != nil {
				panic(err)
			}
			result[left] += count
			result[right] += count
			continue
		}

		// Rule 3: Multiply by 2024
		result[stone*2024] += count
	}
	return result
}
