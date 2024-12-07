package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func findMissingRunes(line string) ([]rune, bool) {
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
				return []rune{}, false
			}
		}
	}
	reverseAny(stack)
	return stack, true
}

func getMissingScore(character rune) int {
	scores := []int{1, 2, 3, 4}
	i := strings.IndexRune(startRunes, character)
	return scores[i]
}

func Part2() {
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

	var scores []int
	for _, line := range lines {
		missing, incomplete := findMissingRunes(line)
		if incomplete {
			score := 0
			for _, r := range missing {
				score *= 5
				score += getMissingScore(r)
			}
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)

	fmt.Println(scores[len(scores)/2])
}
