package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	north = iota
	east  = iota
	south = iota
	west  = iota
)

type Point struct {
	x, y   int
	dx, dy int
	score  int
}

func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(f)
	var grid [][]rune
	var score [][]int
	for scan.Scan() {
		line := scan.Text()

		var runeRow []rune
		var scoreRow []int
		for _, r := range line {
			runeRow = append(runeRow, r)
			scoreRow = append(scoreRow, -1)
		}
		grid = append(grid, runeRow)
		score = append(score, scoreRow)
	}

	// fmt.Println(grid)
	// fmt.Println(score)

	startX, startY := FindRune(grid, 'S')
	score[startY][startX] = 0
	endX, endY := FindRune(grid, 'E')

	queue := []Point{
		{x: startX, y: startY, dx: 1, dy: 0, score: 0},
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if grid[curr.y][curr.x] == 'E' {
			continue
		}

		next := Point{
			x:     curr.x + curr.dx,
			y:     curr.y + curr.dy,
			dx:    curr.dx,
			dy:    curr.dy,
			score: curr.score + 1,
		}
		cwx, cwy := RotateClockwise(curr.dx, curr.dy)
		cw := Point{
			x:     curr.x + cwx,
			y:     curr.y + cwy,
			dx:    cwx,
			dy:    cwy,
			score: curr.score + 1001,
		}
		ccwx, ccwy := RotateCounterClockwise(curr.dx, curr.dy)
		ccw := Point{
			x:     curr.x + ccwx,
			y:     curr.y + ccwy,
			dx:    ccwx,
			dy:    ccwy,
			score: curr.score + 1001,
		}

		for _, p := range []Point{next, cw, ccw} {
			if grid[p.y][p.x] == '#' {
				continue
			}
			if score[p.y][p.x] == -1 {
				score[p.y][p.x] = p.score
				queue = append(queue, p)
			}
			if score[p.y][p.x] > p.score {
				score[p.y][p.x] = p.score
				queue = append(queue, p)
			}
		}
	}

	fmt.Println(score[endY][endX])
}

func RotateClockwise(x, y int) (int, int) {
	if x == 0 && y == 1 {
		return 1, 0
	} else if x == 1 && y == 0 {
		return 0, -1
	} else if x == 0 && y == -1 {
		return -1, 0
	} else if x == -1 && y == 0 {
		return 0, 1
	}
	panic("Bad")
}

func RotateCounterClockwise(x, y int) (int, int) {
	if x == 0 && y == 1 {
		return -1, 0
	} else if x == -1 && y == 0 {
		return 0, -1
	} else if x == 0 && y == -1 {
		return 1, 0
	} else if x == 1 && y == 0 {
		return 0, 1
	}
	panic("Bad")
}

func PrintScore(score [][]int) {
	maxWidth := 0
	for _, row := range score {
		for _, num := range row {
			width := len(fmt.Sprintf("%d", num))
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	fmt.Println()

	for _, row := range score {
		for _, num := range row {
			fmt.Printf("|%5d ", num)
		}
		fmt.Println()
	}
}

func FindRune(grid [][]rune, r rune) (x, y int) {
	for y, row := range grid {
		for x, cell := range row {
			if cell == r {
				return x, y
			}
		}
	}
	return -1, -1
}
