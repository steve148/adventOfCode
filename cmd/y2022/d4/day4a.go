package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse_line(s string) (int, int, int, int) {
	elves := strings.Split(s, ",")
	a := strings.Split(elves[0], "-")
	a1, _ := strconv.Atoi(a[0])
	a2, _ := strconv.Atoi(a[1])
	b := strings.Split(elves[1], "-")
	b1, _ := strconv.Atoi(b[0])
	b2, _ := strconv.Atoi(b[1])
	return a1, a2, b1, b2

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

	var count int
	for _, line := range lines {
		a1, a2, b1, b2 := parse_line(line)

		fmt.Println(a1, a2, b1, b2)

		if (a1 <= b1 && a2 >= b2) || (a1 >= b1 && a2 <= b2) {
			count += 1
		}
	}
	fmt.Println(count)
}
