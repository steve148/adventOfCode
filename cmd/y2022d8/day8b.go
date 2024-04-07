package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type key struct {
	row    int
	column int
}

func calculate_left_score(trees [][]int, row int, column int) int {
	var score int
	for i := column - 1; i >= 0; i-- {
		score++
		if i == 0 || trees[row][i] >= trees[row][column] {
			break
		}
	}
	return score
}

func calculate_right_score(trees [][]int, row int, column int) int {
	var score int
	width := len(trees[0])
	for i := column + 1; i <= width; i++ {
		score++
		if i == width-1 || trees[row][i] >= trees[row][column] {
			break
		}
	}
	return score
}

func calculate_up_score(trees [][]int, row int, column int) int {
	var score int
	for i := row - 1; i >= 0; i-- {
		score++
		if i == 0 || trees[i][column] >= trees[row][column] {
			break
		}
	}
	return score
}

func calculate_down_score(trees [][]int, row int, column int) int {
	var score int
	height := len(trees)
	for i := row + 1; i <= height-1; i++ {
		score++
		if i == height-1 || trees[i][column] >= trees[row][column] {
			break
		}
	}
	return score
}

func calculate_viewing_score(trees [][]int, row int, column int) int {
	left := calculate_left_score(trees, row, column)
	right := calculate_right_score(trees, row, column)
	up := calculate_up_score(trees, row, column)
	down := calculate_down_score(trees, row, column)
	return left * right * up * down
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

	var trees [][]int

	for _, line := range lines {
		var row []int
		for _, s := range strings.Split(line, "") {
			i, _ := strconv.Atoi(s)
			row = append(row, i)
		}
		trees = append(trees, row)
	}

	width := len(trees[0])
	height := len(trees)

	ideal := -1

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			score := calculate_viewing_score(trees, i, j)
			fmt.Println(i, j, score)
			if score > ideal {
				ideal = score
			}
		}
	}

	fmt.Println(ideal)
}
