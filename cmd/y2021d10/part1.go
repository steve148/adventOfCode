package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findFirstIllegalCharacter(line string) (rune, bool) {
	var stack []rune

	for _, r := range line {
		if strings.ContainsRune(startRunes, r) {
			stack = append(stack, r)
		}

		if strings.ContainsRune(endRunes, r) {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i := strings.IndexRune(startRunes, popped)
			expected := endRunes[i]

			if r != rune(expected) {
				return r, true
			}
		}
	}
	return 'a', false
}

func getIllegalScore(character rune) int {
	scores := []int{3, 57, 1197, 25137}
	i := strings.IndexRune(endRunes, character)
	return scores[i]
}

func Part1() {
	file, err := os.Open("./day10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	score := 0
	for _, line := range lines {
		character, illegalFound := findFirstIllegalCharacter(line)
		if illegalFound {
			fmt.Println(line, string(character), illegalFound)
			score += getIllegalScore(character)
		}
	}

	fmt.Println(score)
}
