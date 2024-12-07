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
		isSafe := false

		fmt.Println(report)
		for popIndex := 0; popIndex < len(report); popIndex++ {
			poppedReport := append([]int{}, report[:popIndex]...)
			poppedReport = append(poppedReport, report[popIndex+1:]...)
			fmt.Println(poppedReport)
			popppedRowSafe := true
			for i := 0; i < len(poppedReport)-1; i++ {
				curr := poppedReport[i]
				next := poppedReport[i+1]

				if increasing && (next < curr+1 || next > curr+3) {
					popppedRowSafe = false
					break

				}
				if !increasing && (next > curr-1 || next < curr-3) {
					popppedRowSafe = false
					break
				}
			}

			if popppedRowSafe {
				isSafe = true
				break
			}
		}

		if isSafe {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}
