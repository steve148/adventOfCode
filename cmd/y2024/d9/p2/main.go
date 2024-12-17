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

	i := len(blocks) - 1
	for i >= 0 {
		if blocks[i] == "." {
			i--
		}

		count := 0
		for j := i; j >= 0; j-- {
			if blocks[i] != blocks[j] {
				break
			}
			count++
		}

		j := FindFreeBlock(blocks, count, i)
		if j != -1 {
			for k := j; k < j+count; k++ {
				blocks[k] = blocks[i]
			}
			for k := i - count + 1; k <= i; k++ {
				blocks[k] = "."
			}
		}

		i -= count
	}

	var checksum int
	for i, s := range blocks {
		if s == "." {
			continue
		}
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		checksum += num * i
	}

	fmt.Println(checksum)
}

func FindFreeBlock(blocks []string, num int, end int) int {
	curr := num
	for i := 0; i < end; i++ {
		if blocks[i] == "." {
			curr--
		} else {
			curr = num
		}

		if curr == 0 {
			return i - num + 1
		}
	}
	return -1
}
