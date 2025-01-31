package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(readFile)
	scan.Split(bufio.ScanLines)

	var reports [][]int

	regexNum := regexp.MustCompile(`(\d+)`)

	for scan.Scan() {
		line := scan.Text()
		matches := regexNum.FindAllString(line, -1)

		var report []int
		for _, m := range matches {
			level, err := strconv.Atoi(m)
			if err != nil {
				panic(err)
			}

			report = append(report, level)
		}

		reports = append(reports, report)
	}

	var increasing bool
	var safeCount int
	for _, report := range reports {
		increasing = report[0] < report[1]
		isSafe := true

		for i := 0; i < len(report)-1; i++ {
			curr := report[i]
			next := report[i+1]

			fmt.Println(report[i], report[i+1], increasing)
			if increasing && (next < curr+1 || next > curr+3) {
				isSafe = false
				break

			}
			if !increasing && (next > curr-1 || next < curr-3) {
				isSafe = false
				break
			}
		}
		fmt.Println(report, isSafe)
		if isSafe {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}
