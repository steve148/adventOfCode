package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)
	input = strings.TrimSpace(input)

	productRanges := strings.Split(input, ",")

	sillySum := 0
	for _, productRange := range productRanges {
		s := strings.Split(productRange, "-")
		startStr, endStr := s[0], s[1]
		startNum, _ := strconv.Atoi(startStr)
		endNum, _ := strconv.Atoi(endStr)
		for num := startNum; num <= endNum; num++ {
			numStr := strconv.Itoa(num)
			if isSilly(numStr) {
				fmt.Println(num)
				sillySum += num
			}
		}
	}

	fmt.Println(sillySum)

}

func isSilly(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	firstHalf := s[0 : len(s)/2]
	secondHalf := s[len(s)/2:]

	if firstHalf != secondHalf {
		return false
	}

	return true
}
