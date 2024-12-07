package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items          []int
	operation      string
	divisible_test int
	true_index     int
	false_index    int
}

func apply_operation(monkey *monkey, item int) int {
	split := strings.Split(monkey.operation, " ")

	left := split[0]
	operand := split[1]
	right := split[2]

	if operand == "+" {
		sum := 0
		if left == "old" {
			sum += item
		} else {
			i, _ := strconv.Atoi(left)
			sum += i
		}

		if right == "old" {
			sum += item
		} else {
			i, _ := strconv.Atoi(right)
			sum += i
		}
		return sum
	} else if operand == "*" {
		product := 1
		if left == "old" {
			product *= item
		} else {
			i, _ := strconv.Atoi(left)
			product *= i
		}

		if right == "old" {
			product *= item
		} else {
			i, _ := strconv.Atoi(right)
			product *= i
		}
		return product
	} else {
		log.Fatal("Invalid operand ", operand)
	}

	return item
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

	int_regex := regexp.MustCompile(`\d+`)
	item_limit := 1
	var monkeys []*monkey
	var current_monkey *monkey

	for i, line := range lines {
		switch i % 7 {
		case 0:
			current_monkey = &monkey{}
			monkeys = append(monkeys, current_monkey)
		case 1:
			items := int_regex.FindAllString(line, -1)
			for _, item := range items {
				i, _ := strconv.Atoi(item)
				current_monkey.items = append(current_monkey.items, i)
			}
		case 2:
			// Fetch the RHS of equation for later
			current_monkey.operation = line[19:]
		case 3:
			s := int_regex.FindString(line)
			i, _ := strconv.Atoi(s)
			current_monkey.divisible_test = i
			item_limit *= i
		case 4:
			s := int_regex.FindString(line)
			i, _ := strconv.Atoi(s)
			current_monkey.true_index = i
		case 5:
			s := int_regex.FindString(line)
			i, _ := strconv.Atoi(s)
			current_monkey.false_index = i
		}
	}

	var inspection_count []int
	for i := 0; i < len(monkeys); i++ {
		inspection_count = append(inspection_count, 0)
	}

	num_rounds := 10000
	for round := 0; round < num_rounds; round++ {
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]

			for len(monkey.items) > 0 {
				// Fetch and pop item from slice.
				item := monkey.items[0]
				monkey.items = monkey.items[1:]

				// Increment view count.
				inspection_count[i]++

				// Calculate new item value.
				new_item := apply_operation(monkey, item) % item_limit

				// Send item to next applicable monkey.
				if new_item%monkey.divisible_test == 0 {
					monkeys[monkey.true_index].items = append(monkeys[monkey.true_index].items, new_item)
				} else {
					monkeys[monkey.false_index].items = append(monkeys[monkey.false_index].items, new_item)
				}
			}
		}
	}

	sort.Ints(inspection_count)
	highest := inspection_count[len(inspection_count)-1]
	second_highest := inspection_count[len(inspection_count)-2]

	fmt.Println(highest*second_highest, highest, second_highest)
}
