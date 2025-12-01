package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Computer struct {
	registers [3]int
	program   []int
	curr      int
	output    []int
}

func (c *Computer) IsComplete() bool {
	return c.curr >= len(c.program)
}

func (c Computer) String() string {
	return fmt.Sprintf("Registers: %v, Program: %v, Current: %d, Out: %v", c.registers, c.program, c.curr, c.output)
}

func (c *Computer) ComboVal() int {
	i := c.program[c.curr+1]
	if i <= 3 {
		return i
	}
	if i <= 6 {
		return c.registers[i-4]
	}
	panic("Invalid combo")
}

func (c *Computer) adv() {
	combo := c.ComboVal()
	num := c.registers[0]
	den := powInt(2, combo)
	c.registers[0] = num / den
	c.curr += 2
}

func (c *Computer) bxl() {
	b := c.registers[1]
	lit := c.program[c.curr+1]
	xor := b ^ lit
	c.registers[1] = xor
	c.curr += 2
}

func (c *Computer) bst() {
	combo := c.ComboVal()
	mod := combo % 8
	c.registers[1] = mod
	c.curr += 2
}

func (c *Computer) jnz() {
	if c.registers[0] != 0 {
		lit := c.program[c.curr+1]
		c.curr = lit - 2
	}
	c.curr += 2
}

func (c *Computer) bxc() {
	num1 := c.registers[1]
	num2 := c.registers[2]
	res := num1 ^ num2
	c.registers[1] = res
	c.curr += 2
}

func (c *Computer) out() {
	combo := c.ComboVal()
	mod := combo % 8
	c.output = append(c.output, mod)
	c.curr += 2
}

func (c *Computer) bdv() {
	combo := c.ComboVal()
	num := c.registers[0]
	den := powInt(2, combo)
	c.registers[1] = num / den
	c.curr += 2
}

func (c *Computer) cdv() {
	combo := c.ComboVal()
	num := c.registers[0]
	den := powInt(2, combo)
	c.registers[2] = num / den
	c.curr += 2
}

func (c *Computer) Result() string {
	var outputStrings []string
	for _, o := range c.output {
		outputStrings = append(outputStrings, strconv.Itoa(o))
	}
	return strings.Join(outputStrings, "")
}

func (c *Computer) IsOutputCopy() bool {
	if len(c.program) != len(c.output) {
		return false
	}

	for i := 0; i < len(c.program); i++ {
		if c.program[i] != c.output[i] {
			return false
		}
	}

	return true
}

func (c *Computer) Tick() {
	switch c.program[c.curr] {
	case 0:
		c.adv()
	case 1:
		c.bxl()
	case 2:
		c.bst()
	case 3:
		c.jnz()
	case 4:
		c.bxc()
	case 5:
		c.out()
	case 6:
		c.bdv()
	case 7:
		c.cdv()
	default:
		panic("Invalid operator")
	}
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
func main() {
	f, err := os.Open("data.txt")
	if err != nil {
		panic(err)
	}

	regexA := regexp.MustCompile(`Register A: (\d+)`)
	regexB := regexp.MustCompile(`Register B: (\d+)`)
	regexC := regexp.MustCompile(`Register C: (\d+)`)
	regexP := regexp.MustCompile(`Program`)
	numRegex := regexp.MustCompile(`(\d+)`)

	var registers [3]int
	var program []int
	scan := bufio.NewScanner(f)
	for scan.Scan() {
		line := scan.Text()

		matchA := regexA.FindStringSubmatch(line)
		if len(matchA) > 0 {
			num, err := strconv.Atoi(matchA[1])
			if err != nil {
				panic(err)
			}
			registers[0] = num
		}
		matchB := regexB.FindStringSubmatch(line)
		if len(matchB) > 0 {
			num, err := strconv.Atoi(matchB[1])
			if err != nil {
				panic(err)
			}
			registers[1] = num
		}
		matchC := regexC.FindStringSubmatch(line)
		if len(matchC) > 0 {
			num, err := strconv.Atoi(matchC[1])
			if err != nil {
				panic(err)
			}
			registers[2] = num
		}
		matchP := regexP.FindStringSubmatch(line)
		numbers := numRegex.FindAllString(line, -1)
		if len(matchP) > 0 {
			for _, number := range numbers {
				num, err := strconv.Atoi(number)
				if err != nil {
					panic(err)
				}
				program = append(program, num)
			}
		}
	}

	for i := 0; i < 1000000000000; i++ {
		var regCopy [3]int
		_ = copy(registers[:], regCopy[:])
		regCopy[0] = i

		comp := Computer{
			registers: regCopy,
			program:   program,
			curr:      0,
			output:    []int{},
		}
		for !comp.IsComplete() {
			comp.Tick()
		}

		if comp.IsOutputCopy() {
			fmt.Println(i)
			break
		}
	}
}
