package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	b, err := os.ReadFile("./data.txt")
	if err != nil {
		panic(err)
	}

	s := string(b)

	r := regexp.MustCompile(`^mul\((\d+),(\d+)\)`)

	sum := 0
	enable := true
	for i := range s {

		if len(s)-i-7 == 0 {
			break
		}

		if s[i:i+4] == "do()" {
			enable = true
			fmt.Println(enable)
		}
		if s[i:i+7] == "don't()" {
			enable = false
			fmt.Println(enable)
		}
		matches := r.FindStringSubmatch(s[i:])
		if len(matches) == 3 {
			left, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			right, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}

			if enable {
				sum += left * right
			}
		}
	}
	fmt.Println(sum)
}
