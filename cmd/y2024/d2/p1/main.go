package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	readFile, err := os.Open("data/2024/d2/data.txt")
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

		for i := 0; i < len(report); i++ {
			if i > 0 {
				if increasing && report[i-1] < report[i] && report[i-1]+4 > report[i] {
					isSafe = false
					fmt.Println(1)
					break
				}
				if !increasing && report[i-1] > report[i] && report[i-1]-4 < report[i] {
					isSafe = false
					fmt.Println(2)
					break
				}
			}
			if i < len(report)-1 {
				if increasing && report[i] < report[i+1] {
					isSafe = false
					fmt.Println(3)
					break
				}
				if !increasing && report[i] > report[i+1] {
					isSafe = false
					fmt.Println(4)
					break
				}
			}
		}
		fmt.Println(report, isSafe)
		if isSafe {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}
