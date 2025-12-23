package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	var batteryRows []string
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		batteryRows = append(batteryRows, line)
	}

	sumJoltage := 0
	for _, row := range batteryRows {
		var joltageStr string

		minIndex := 0
		maxIndex := len(row) - 12

		for i := 0; i < 12; i++ {
			newMinIndex := minIndex
			maxBattery := "0"

			for j := minIndex; j <= maxIndex; j++ {
				newBattery := string(row[j])

				if newBattery > maxBattery {
					maxBattery = newBattery
					newMinIndex = j
				}
			}

			joltageStr += maxBattery
			minIndex = newMinIndex + 1
			maxIndex += 1
		}

		joltage, _ := strconv.Atoi(joltageStr)
		sumJoltage += joltage
	}

	fmt.Println(sumJoltage)
}
