package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
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

	for i := 0; i < 25; i++ {
		stones = Blink(stones)
	}
	fmt.Println(len(stones))
}

func Blink(stones []int) []int {
	var result []int
	for _, stone := range stones {

		if stone == 0 {
			// Rule 1: Swap 0 for 1
			result = append(result, 1)
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
			result = append(result, left, right)
			continue
		}

		// Rule 3: Multiply by 2024
		result = append(result, stone*2024)
	}
	return result
}
