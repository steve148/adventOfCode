package day4

type BingoBoard struct {
	numbers  [][]int
	marked   [5][5]bool
	complete bool
}

func (b *BingoBoard) isComplete() bool {
	return b.complete
}

func (b *BingoBoard) updateBoard(x int) {
	for i := 0; i < len(b.numbers); i++ {
		for j := 0; j < len(b.numbers[i]); j++ {
			if b.numbers[i][j] == x {
				b.marked[i][j] = true
			}
		}
	}
}

func (b *BingoBoard) isBingo() bool {
	// Do any rows have a full columns
	for i := 0; i < len(b.marked); i++ {
		if isBingo(b.marked[i][:]) {
			return true
		}
	}
	// Do any columns have a match
	for i := 0; i < len(b.marked); i++ {
		var column []bool
		for _, row := range b.marked {
			column = append(column, row[i])
		}
		if isBingo(column) {
			return true
		}
	}

	return false
}

func isBingo(b []bool) bool {
	for _, v := range b {
		if !v {
			return false
		}
	}
	return true
}

func (b *BingoBoard) sumUnmarked() int {
	sum := 0
	for i := 0; i < len(b.marked); i++ {
		for j := 0; j < len(b.marked[i]); j++ {
			if !b.marked[i][j] {
				sum += b.numbers[i][j]
			}
		}
	}
	return sum
}
