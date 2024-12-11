package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Page struct {
	id     int
	before map[int]bool
	after  map[int]bool
}

func NewPage(id int) *Page {
	return &Page{id: id, before: make(map[int]bool), after: make(map[int]bool)}
}

func (p Page) HasAfter(id int) bool {
	_, ok := p.after[id]
	return ok
}

func (p Page) HasBefore(id int) bool {
	_, ok := p.before[id]
	return ok
}

var pageMap = make(map[int]*Page)
var updates [][]int

func main() {
	readFile, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}

	scan := bufio.NewScanner(readFile)
	scan.Split(bufio.ScanLines)

	scanningRules := true
	scanningPages := false

	rePage := regexp.MustCompile(`(\d+)`)

	for scan.Scan() {
		line := scan.Text()

		if line == "" {
			scanningRules = false
			scanningPages = true

		} else if scanningRules {
			matches := rePage.FindAllString(line, -1)
			leftStr := matches[0]
			rightStr := matches[1]
			leftID, err := strconv.Atoi(leftStr)
			if err != nil {
				panic(err)
			}
			rightID, err := strconv.Atoi(rightStr)
			if err != nil {
				panic(err)
			}

			if _, ok := pageMap[leftID]; !ok {
				pageMap[leftID] = NewPage(leftID)
			}
			if _, ok := pageMap[rightID]; !ok {
				pageMap[rightID] = NewPage(rightID)
			}

			pageMap[leftID].after[rightID] = true
			pageMap[rightID].before[leftID] = true
		} else if scanningPages {
			matches := rePage.FindAllString(line, -1)
			var update []int
			for _, match := range matches {
				pageID, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				update = append(update, pageID)

			}
			updates = append(updates, update)
		} else {
			panic("never")
		}
	}

	validSum := 0
	for _, update := range updates {
		updateValid := true
		for i, pageID := range update {
			valid := true
			page := pageMap[pageID]

			for j := i; j < len(update); j++ {
				if page.HasBefore(update[j]) {
					valid = false
					break
				}
			}

			if !valid {
				updateValid = false
				break
			}
		}
		if updateValid {
			middlePage := update[len(update)/2]
			validSum += middlePage
		}

	}
	fmt.Println(validSum)
}
