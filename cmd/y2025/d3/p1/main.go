package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	sumJoltage := 0

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		chars := strings.Split(line, "")

		var nums []int
		for _, s := range chars {
			i, _ := strconv.Atoi(s)
			nums = append(nums, i)
		}

		joltage := 0
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				newJoltage := nums[i]*10 + nums[j]

				if newJoltage > joltage {
					joltage = newJoltage
				}
			}
		}

		sumJoltage += joltage
	}

	fmt.Println(sumJoltage)
}
