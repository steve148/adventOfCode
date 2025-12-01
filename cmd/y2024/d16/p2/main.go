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
	x, y int
}

type QueueElem struct {
	p      Point
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

	queue := []QueueElem{
		{p: Point{x: startX, y: startY}, dx: 1, dy: 0, score: 0},
	}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if grid[curr.p.y][curr.p.x] == 'E' {
			continue
		}

		next := QueueElem{
			p: Point{
				x: curr.p.x + curr.dx,
				y: curr.p.y + curr.dy,
			},
			dx:    curr.dx,
			dy:    curr.dy,
			score: curr.score + 1,
		}
		cwx, cwy := RotateClockwise(curr.dx, curr.dy)
		cw := QueueElem{
			p: Point{
				x: curr.p.x + cwx,
				y: curr.p.y + cwy,
			},
			dx:    cwx,
			dy:    cwy,
			score: curr.score + 1001,
		}
		ccwx, ccwy := RotateCounterClockwise(curr.dx, curr.dy)
		ccw := QueueElem{
			p: Point{
				x: curr.p.x + ccwx,
				y: curr.p.y + ccwy,
			},
			dx:    ccwx,
			dy:    ccwy,
			score: curr.score + 1001,
		}

		for _, qe := range []QueueElem{next, cw, ccw} {
			p := qe.p
			if grid[p.y][p.x] == '#' {
				continue
			}
			if score[p.y][p.x] == -1 {
				score[p.y][p.x] = qe.score
				queue = append(queue, qe)
			}
			if score[p.y][p.x] > qe.score {
				score[p.y][p.x] = qe.score
				queue = append(queue, qe)
			}
		}
	}

	PrintScore(score)

	seen := make(map[Point]bool)
	seen[Point{x: endX, y: endY}] = true
	seen[Point{x: startX, y: startY}] = true
	pQueue := []Point{
		{x: endX, y: endY},
	}
	for len(pQueue) > 0 {
		curr := pQueue[0]
		pQueue = pQueue[1:]

		left := Point{x: curr.x - 1, y: curr.y}
		right := Point{x: curr.x + 1, y: curr.y}
		up := Point{x: curr.x, y: curr.y - 1}
		down := Point{x: curr.x, y: curr.y + 1}

		currScore := score[curr.y][curr.x]
		for _, p := range []Point{left, right, up, down} {
			pScore := score[p.y][p.x]
			if pScore == -1 {
				continue
			}

			if seen[p] {
				continue
			}

			if pScore-1000 < currScore {
				seen[p] = true
				pQueue = append(pQueue, p)
			}
		}
	}

	PrintGridSeen(grid, seen)
	fmt.Println(len(seen))
}

func PrintGridSeen(grid [][]rune, seen map[Point]bool) {
	for y, row := range grid {
		for x, cell := range row {
			if seen[Point{x: x, y: y}] {
				fmt.Print("O")
			} else {
				fmt.Print(string(cell))
			}
		}
		fmt.Println()
	}

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
