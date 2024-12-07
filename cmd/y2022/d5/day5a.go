package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func parse_line(s string) []int {
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllStringSubmatch(s, -1)
	var result []int
	for _, m := range matches {
		i, _ := strconv.Atoi(m[0])
		result = append(result, i)
	}
	return result
}

func find_header_index(lines []string) int {
	for i, line := range lines {
		if found, _ := regexp.MatchString(`\d+`, line); found {
			return i
		}
	}
	return -1
}

func init_stacks(lines []string, header_index int) []Stack {
	var stacks []Stack

	line := lines[0]
	num_stacks := (len(line) + 1) / 4

	for i := 0; i < num_stacks; i++ {
		var s Stack
		stacks = append(stacks, s)
	}

	for line_index := header_index - 1; line_index >= 0; line_index-- {
		for i := 0; i < num_stacks; i++ {
			start := i * 4
			end := i*4 + 3
			r := regexp.MustCompile(`[A-Z]`)
			s := r.FindString(lines[line_index][start:end])
			if s != "" {
				stacks[i].Push(s)
			}
		}

	}

	return stacks
}

func parse_step(line string) []int {
	r := regexp.MustCompile(`\d+`)
	matches := r.FindAllStringSubmatch(line, -1)
	var result []int
	for _, m := range matches {
		i, _ := strconv.Atoi(m[0])
		result = append(result, i)
	}
	return result
}

func main() {
	file, err := os.Open("./input-sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	header_index := find_header_index(lines)

	stacks := init_stacks(lines, header_index)

	for i := header_index + 2; i < len(lines); i++ {

		line := lines[i]
		step := parse_step(line)
		num, from, to := step[0], step[1], step[2]

		from_stack := &stacks[from-1]
		to_stack := &stacks[to-1]

		for i := 0; i < num; i++ {
			item, _ := from_stack.Pop()
			to_stack.Push(item)
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
	fmt.Println()

}
