package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Packet struct {
	value       int
	sub_packets []*Packet
	parent      *Packet
}

func (p *Packet) is_leaf() bool {
	return len(p.sub_packets) == 0
}

func NewPacket(s string) *Packet {
	var stack []*Packet
	var result *Packet
	var runes []rune

	for _, r := range s {
		switch r {
		case '[':
			new := &Packet{value: -1}
			if len(stack) != 0 {
				parent := stack[len(stack)-1]
				new.parent = parent
				parent.sub_packets = append(parent.sub_packets, new)
			}
			stack = append(stack, new)
		case ']':
			if len(runes) != 0 {
				i, _ := strconv.Atoi(string(runes))
				parent := stack[len(stack)-1]
				new := &Packet{value: i, parent: parent}
				parent.sub_packets = append(parent.sub_packets, new)
				runes = nil
			}
			result = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		case ',':
			i, _ := strconv.Atoi(string(runes))
			parent := stack[len(stack)-1]
			new := &Packet{value: i, parent: parent}
			parent.sub_packets = append(parent.sub_packets, new)
			runes = nil
		default:
			runes = append(runes, r)
		}
	}

	return result
}

func print_packet(p *Packet, indent int) {
	prefix := strings.Repeat("-", indent)

	if prefix != "" {
		prefix += " "
	}

	fmt.Print(prefix, "Packet ", p.value)
	fmt.Println()

	for _, sub_packet := range p.sub_packets {
		print_packet(sub_packet, indent+2)
	}
}

func are_packets_ordered(l, r *Packet) int {
	if l.is_leaf() && r.is_leaf() {
		if l.value < r.value {
			return 1
		}
		if l.value > r.value {
			return -1
		}
		return 0
	}

	if !l.is_leaf() && r.is_leaf() {
		return are_packets_ordered(l, &Packet{value: -1, sub_packets: []*Packet{r}})
	}

	if l.is_leaf() && !r.is_leaf() {
		return are_packets_ordered(&Packet{value: -1, sub_packets: []*Packet{l}}, r)
	}

	if !l.is_leaf() && !r.is_leaf() {
		var i int
		for i = 0; i < len(r.sub_packets) && i < len(l.sub_packets); i++ {
			sub_packets_ordered := are_packets_ordered(l.sub_packets[i], r.sub_packets[i])
			if sub_packets_ordered != 0 {
				return sub_packets_ordered
			}
		}

		if i < len(l.sub_packets) {
			return -1
		} else if i < len(r.sub_packets) {
			return 1
		}
	}
	return 0
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var packets []*Packet

	packets = append(packets, NewPacket("[[2]]"))
	packets = append(packets, NewPacket("[[6]]"))

	for _, line := range lines {
		if line != "" {
			packets = append(packets, NewPacket(line))
		}
	}

	p1 := packets[0]
	p2 := packets[1]

	sort.Slice(packets, func(i, j int) bool {
		return are_packets_ordered(packets[i], packets[j]) == 1
	})

	product := 1
	for i := 0; i < len(packets); i++ {
		if packets[i] == p1 || packets[i] == p2 {
			product *= (i + 1)
		}
	}

	fmt.Println(product)
}
