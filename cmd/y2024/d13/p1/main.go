package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	b    string
	x, y int
}

type Machine struct {
	x, y    int
	buttonA Button
	buttonB Button
}

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	buttonRegex := regexp.MustCompile(`Button (A|B): X\+(\d+), Y\+(\d+)`)
	prizeRegex := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	machines := []Machine{}
	machine := &Machine{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "Button A") {
			m := buttonRegex.FindAllStringSubmatch(line, -1)
			x, err := strconv.Atoi(m[0][2])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(m[0][3])
			if err != nil {
				panic(err)
			}
			machine.buttonA = Button{m[0][1], x, y}
		} else if strings.HasPrefix(line, "Button B") {
			m := buttonRegex.FindAllStringSubmatch(line, -1)
			x, err := strconv.Atoi(m[0][2])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(m[0][3])
			if err != nil {
				panic(err)
			}
			machine.buttonB = Button{m[0][1], x, y}
		} else if strings.HasPrefix(line, "Prize") {
			m := prizeRegex.FindAllStringSubmatch(line, -1)
			x, err := strconv.Atoi(m[0][1])
			if err != nil {
				panic(err)
			}
			y, err := strconv.Atoi(m[0][2])
			if err != nil {
				panic(err)
			}
			machine.x = x
			machine.y = y
		} else {
			machines = append(machines, *machine)
			machine = &Machine{}
		}
	}

	maxPress := 100
	infCost := 999999999999
	totalCost := 0
	for _, m := range machines {
		minCost := infCost
		for i := 0; i < maxPress; i++ {
			for j := 0; j < maxPress; j++ {
				x := i*m.buttonA.x + j*m.buttonB.x
				y := i*m.buttonA.y + j*m.buttonB.y
				cost := i*3 + j
				if x == m.x && y == m.y && cost < minCost {
					minCost = cost
				}
			}
		}
		if minCost < infCost {
			totalCost += minCost
		}
	}

	fmt.Println(totalCost)
}
