package day14

import "fmt"

func parseInput(lines []string) (map[string]int, map[string]string, map[string]int) {
	pairs := make(map[string]int)
	rules := make(map[string]string)

	polymer := lines[0]
	counts := map[string]int{polymer[0:1]: 1}
	for i := 0; i < len(polymer)-1; i++ {
		pairs[polymer[i:i+2]]++
		counts[polymer[i+1:i+2]]++
	}

	for _, line := range lines[2:] {
		var pattern string
		var insert string
		fmt.Sscanf(line, "%s -> %s", &pattern, &insert)
		rules[pattern] = insert
	}

	return pairs, rules, counts
}
