package day6

import (
	"fmt"
	"strconv"
	"strings"
)

type State [9]int

func parseInput(lines []string) State {
	var initial_state State
	for _, s := range strings.Split(lines[0], ",") {
		age, _ := strconv.Atoi(s)
		initial_state[age] += 1
	}
	return initial_state
}

func updateState(state State) State {
	var new_state State
	new_state[0] = state[1]
	new_state[1] = state[2]
	new_state[2] = state[3]
	new_state[3] = state[4]
	new_state[4] = state[5]
	new_state[5] = state[6]
	new_state[6] = state[7] + state[0]
	new_state[7] = state[8]
	new_state[8] = state[0]
	return new_state
}

func sumState(state State) int {
	result := 0
	for _, x := range state {
		result += x
	}
	return result
}

func printState(state State, day int) {
	fmt.Printf("After %v days: %v\n", day, state)
}
