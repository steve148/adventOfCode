package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func playToWin(numbers []int, boards []*BingoBoard) int {
	for _, number := range numbers {
		for _, board := range boards {
			board.updateBoard(number)
			if board.isBingo() {
				return board.sumUnmarked() * number
			}

		}
	}
	log.Fatal("No winner found")
	return 0
}

func Part1() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var chosen_numbers []int
	for _, s := range strings.Split(lines[0], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		chosen_numbers = append(chosen_numbers, i)
	}

	var boards []*BingoBoard

	var numbers [][]int
	for _, line := range lines[1:] {
		trimmed := strings.TrimSpace(line)
		row_strings := strings.Fields(trimmed)
		var row_numbers []int
		for _, s := range row_strings {
			i, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			row_numbers = append(row_numbers, i)
		}

		if line != "" {
			numbers = append(numbers, row_numbers)
		}

		if len(numbers) == 5 {
			var marked [5][5]bool
			board := BingoBoard{numbers: numbers, marked: marked, complete: false}
			boards = append(boards, &board)
			numbers = nil
		}
	}

	score := playToWin(chosen_numbers, boards)

	fmt.Println(score)
}
