package main

import (
	"bufio"
	"fmt"
	"os"
)

type Robot struct {
	x, y   int
	vX, vY int
}

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var robots []*Robot

	width, height := 101, 103
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, vX, vY int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vX, &vY)
		robots = append(robots, &Robot{x: x, y: y, vX: vX, vY: vY})
	}

	PrintRobots(robots)

	for i := 0; i < 1000000; i++ {
		for _, robot := range robots {
			x2, y2 := robot.x+robot.vX, robot.y+robot.vY
			if x2 < 0 {
				x2 += width
			}
			if y2 < 0 {
				y2 += height
			}
			if x2 >= width {
				x2 -= width
			}
			if y2 >= height {
				y2 -= height
			}
			robot.x, robot.y = x2, y2
		}

		if MaybeTree(robots) {
			PrintRobots(robots)
			fmt.Println("Iteration:", i)
			fmt.Println("Press Enter to continue...")
			reader := bufio.NewReader(os.Stdin)
			_, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
		}
	}

	var q1, q2, q3, q4 int
	for _, robot := range robots {
		if robot.x < width/2 && robot.y < height/2 {
			q1++
		} else if robot.x > width/2 && robot.y < height/2 {
			q2++
		} else if robot.x < width/2 && robot.y > height/2 {
			q3++
		} else if robot.x > width/2 && robot.y > height/2 {
			q4++
		}
	}
	fmt.Println(q1, q2, q3, q4, q1*q2*q3*q4)
}

func MaybeTree(robots []*Robot) bool {
	grid := make([][]int, 103)
	for i := range grid {
		grid[i] = make([]int, 103)
	}
	for _, robot := range robots {
		grid[robot.y][robot.x]++
	}
	for y := 0 + 2; y < 103-2; y++ {
		for x := 0 + 2; x < 101-2; x++ {
			if grid[y][x] > 0 && grid[y-1][x-1] > 0 && grid[y-1][x] > 0 && grid[y-1][x+1] > 0 && grid[y-2][x-2] > 0 && grid[y-2][x-1] > 0 && grid[y-2][x] > 0 && grid[y-2][x+1] > 0 && grid[y-2][x+2] > 0 {
				return true
			}
		}
	}
	return false
}

func PrintRobots(robots []*Robot) {
	grid := make([][]int, 103)
	for i := range grid {
		grid[i] = make([]int, 103)
	}
	for _, robot := range robots {
		grid[robot.y][robot.x]++
	}
	for _, row := range grid {
		for _, cell := range row {
			if cell > 0 {
				fmt.Print(cell)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
