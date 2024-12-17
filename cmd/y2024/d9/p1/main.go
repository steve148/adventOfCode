package main

import (
	"bufio"
	"fmt"
	"os"
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

	var blocks []string
	blockCount := 0

	for i, r := range line {
		isFile := i%2 == 0
		blockId := strconv.Itoa(i / 2)

		count, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}

		if isFile {
			blockCount += count
		}

		for j := 0; j < count; j++ {
			if isFile {
				blocks = append(blocks, blockId)
			} else {
				blocks = append(blocks, ".")
			}
		}
	}

	fmt.Println(blocks)

	result := make([]string, blockCount)
	i := 0
	j := len(blocks) - 1
	for i <= j {
		if blocks[i] != "." {
			result[i] = blocks[i]
			i++
		} else if blocks[j] == "." {
			j--
		} else {
			result[i] = blocks[j]
			i++
			j--
		}
	}

	fmt.Println(result)

	var checksum int
	for i, r := range result {
		num, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}

		checksum += num * i
	}

	fmt.Println(checksum)
}
