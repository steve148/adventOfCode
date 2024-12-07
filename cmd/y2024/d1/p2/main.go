package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("data/2024/d1/data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(readFile)
	scan.Split(bufio.ScanLines)

	var left []int
	var right []int

	r := regexp.MustCompile(`(\d+)`)

	for scan.Scan() {
		line := scan.Text()
		matches := r.FindAllString(line, -1)

		left_num, err := strconv.Atoi(matches[0])
		if err != nil {
			panic(err)
		}
		left = append(left, left_num)

		right_num, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		right = append(right, right_num)
	}

	var rightMap map[int]int = make(map[int]int)
	for _, num := range right {
		if _, ok := rightMap[num]; !ok {
			rightMap[num] = 0
		}
		rightMap[num]++
	}

	var num int
	for _, l := range left {
		n, ok := rightMap[l]
		if !ok {
			continue
		}
		num = num + (l * n)
	}
	fmt.Println(num)

	// var diff int
	// for i := 0; i < len(left); i++ {
	// 	if left[i] < right[i] {
	// 		diff = diff + (right[i] - left[i])
	// 	} else {
	// 		diff = diff + (left[i] - right[i])
	// 	}
	// }

	// fmt.Println(diff)
}
