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

	for i := 0; i < 100; i++ {
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
}
