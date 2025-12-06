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
				sillySum += num
			}
		}
	}

	fmt.Println(sillySum)

}

func isSilly(s string) bool {
	length := len(s)

	if length < 2 {
		return false
	}

	for size := 1; size <= length/2; size++ {

		if length%size != 0 {
			continue
		}

		pattern := s[:size]

		isMatch := true
		for i := size; i < length; i += size {
			if s[i:i+size] != pattern {
				isMatch = false
				break
			}
		}

		if isMatch {
			return true
		}
	}

	return false
}
