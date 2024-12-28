package main

import (
	"fmt"
	"os"
	"strings"
)

type Button struct {
	b    string
	x, y float64
}

type Machine struct {
	x, y    float64
	buttonA Button
	buttonB Button
}

func main() {
	input, _ := os.ReadFile("data.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	total := 0

	for i := 0; i < len(lines); i += 4 {
		var aX, aY, bX, bY, pX, pY int
		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &aX, &aY)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &bX, &bY)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &pX, &pY)

		pX += 10000000000000
		pY += 10000000000000

		D, Dx, Dy := aX*bY-bX*aY, pX*bY-bX*pY, aX*pY-pX*aY
		if D != 0 && Dx == (Dx/D)*D && Dy == (Dy/D)*D {
			total += (Dx/D)*3 + (Dy / D)
		}
	}
	fmt.Println(total)
}
