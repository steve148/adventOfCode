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

	current := 50
	count := 0

	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()
		direction := line[0]
		turn, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		fmt.Println(current, count, string(direction), turn)

		if turn > 100 {
			count += turn / 100
			turn %= 100
		}

		if turn == 0 {
			continue
		}

		next := current

		switch direction {
		case 'L':
			next -= turn
		case 'R':
			next += turn
		default:
			panic("Invalid direction")
		}

		if next < 0 {
			if current != 0 {
				count++
			}
			next += 100
		} else if next >= 100 {
			if current != 0 {
				count++
			}
			next -= 100
		} else if next == 0 {
			count++
		}

		current = next

		fmt.Println(current, count)
	}

	fmt.Println(count)
}
