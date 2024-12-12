package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Equation struct {
	result int
	values []int
}

func main() {
	readFile, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(readFile)
	scan.Split(bufio.ScanLines)

	num_re := regexp.MustCompile(`\d+`)

	var equations []Equation
	for scan.Scan() {
		line := scan.Text()

		matches := num_re.FindAllString(line, -1)
		var nums []int
		for _, match := range matches {
			num, err := strconv.Atoi(match)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}

		equation := Equation{
			result: nums[0],
			values: nums[1:],
		}
		equations = append(equations, equation)
	}

	equationCount := 0
	for _, equation := range equations {
		results := []int{equation.values[0]}

		for _, value := range equation.values[1:] {
			var newResults []int
			for _, result := range results {
				newResults = append(newResults, result+value)
				newResults = append(newResults, result*value)

				s := strconv.Itoa(result) + strconv.Itoa(value)
				i, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				newResults = append(newResults, i)
			}
			results = newResults
		}
		for _, results := range results {
			if results == equation.result {
				equationCount += equation.result
				break
			}
		}
	}
	fmt.Println(equationCount)
}
