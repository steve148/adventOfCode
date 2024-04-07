package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func buildDecoder(inputs []string) map[string]string {
	decoder := make(map[string]string)
	wireCount := make(map[rune]int)

	for _, input := range inputs {
		for _, r := range input {
			wireCount[r]++
		}
	}

	for _, input := range inputs {
		sorted := SortString(input)

		switch len(sorted) {
		case 2:
			decoder[sorted] = "1"
		case 3:
			decoder[sorted] = "7"
		case 4:
			decoder[sorted] = "4"
		case 5:
			for _, r := range sorted {
				if wireCount[r] == 4 {
					decoder[sorted] = "2"
				} else if wireCount[r] == 6 {
					decoder[sorted] = "5"
				}
			}
			if _, twoOrFive := decoder[sorted]; !twoOrFive {
				decoder[sorted] = "3"
			}
		case 6:
			wireCountKeyRunes := make(map[int]int)
			for _, r := range sorted {
				wireCountKeyRunes[wireCount[r]]++
			}
			if wireCountKeyRunes[4] == 0 {
				decoder[sorted] = "9"
			} else if wireCountKeyRunes[7] == 1 {
				decoder[sorted] = "0"
			} else {
				decoder[sorted] = "6"
			}
		case 7:
			decoder[sorted] = "8"
		}
	}
	return decoder
}

func Part2() {
	file, err := os.Open("./day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := 0
	for _, line := range lines {

		split := strings.Fields(line)
		inputs := split[:10]
		outputs := split[11:]
		display := ""

		decoder := buildDecoder(inputs)

		for _, code := range outputs {
			value := decoder[SortString(code)]
			display += value
		}
		num, _ := strconv.Atoi(display)
		sum += num
	}
	fmt.Println(sum)
}
