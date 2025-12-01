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

	state := 50
	count := 0

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		direction := line[0]
		turn, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		// fmt.Println(line)

		switch direction {
		case 'L':
			state -= turn % 100
		case 'R':
			state += turn % 100
		default:
			panic("Invalid direction")
		}

		if state > 99 {
			state -= 100
		}
		if state < 0 {
			state += 100
		}

		// fmt.Println(state)

		if state == 0 {
			count++
		}
	}

	fmt.Println(count)
}
