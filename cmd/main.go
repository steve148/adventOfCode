package main

import (
	"os"

	"github.com/steve148/advent-of-code-2021/day10"
	"github.com/steve148/advent-of-code-2021/day11"
	"github.com/steve148/advent-of-code-2021/day12"
	"github.com/steve148/advent-of-code-2021/day13"
	"github.com/steve148/advent-of-code-2021/day14"
	"github.com/steve148/advent-of-code-2021/day4"
	"github.com/steve148/advent-of-code-2021/day5"
	"github.com/steve148/advent-of-code-2021/day6"
	"github.com/steve148/advent-of-code-2021/day7"
	"github.com/steve148/advent-of-code-2021/day8"
	"github.com/steve148/advent-of-code-2021/day9"
)

func main() {
	args := os.Args[1:]
	day := args[0]
	part := args[1]

	if day == "4" && part == "1" {
		day4.Part1()
	} else if day == "4" && part == "2" {
		day4.Part2()
	} else if day == "5" && part == "1" {
		day5.Part1()
	} else if day == "5" && part == "2" {
		day5.Part2()
	} else if day == "6" && part == "1" {
		day6.Part1()
	} else if day == "6" && part == "2" {
		day6.Part2()
	} else if day == "7" && part == "1" {
		day7.Part1()
	} else if day == "7" && part == "2" {
		day7.Part2()
	} else if day == "8" && part == "1" {
		day8.Part1()
	} else if day == "8" && part == "2" {
		day8.Part2()
	} else if day == "9" && part == "1" {
		day9.Part1()
	} else if day == "9" && part == "2" {
		day9.Part2()
	} else if day == "10" && part == "1" {
		day10.Part1()
	} else if day == "10" && part == "2" {
		day10.Part2()
	} else if day == "11" && part == "1" {
		day11.Part1()
	} else if day == "11" && part == "2" {
		day11.Part2()
	} else if day == "12" && part == "1" {
		day12.Part1()
	} else if day == "12" && part == "2" {
		day12.Part2()
	} else if day == "13" && part == "1" {
		day13.Part1()
	} else if day == "13" && part == "2" {
		day13.Part2()
	} else if day == "14" && part == "1" {
		day14.Part1()
	} else if day == "14" && part == "2" {
		day14.Part2()
	}
}
