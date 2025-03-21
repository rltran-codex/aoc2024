package main

import (
	"fmt"
	"io"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Computer struct {
	RegA     uint64
	RegB     uint64
	RegC     uint64
	CurrPtr  uint64
	ProInstr []int
	OutInstr []int
}

func (c *Computer) execProgram() []int {
	for c.CurrPtr = uint64(0); c.CurrPtr < uint64(len(c.ProInstr)); c.CurrPtr += 2 {
		opcode := c.ProInstr[c.CurrPtr]
		operand := c.ProInstr[c.CurrPtr+1]
		c.handleOpcode(opcode, operand)
	}

	return c.OutInstr
}

func (c *Computer) handleOpcode(opcode int, x int) {
	var operand uint64
	switch x {
	case 0, 1, 2, 3:
		operand = uint64(x)
	case 4:
		operand = c.RegA
	case 5:
		operand = c.RegB
	case 6:
		operand = c.RegC
	case 7:
		operand = uint64(x)
	}

	switch opcode {
	case 0: // The adv instruction (opcode 0) division, numerator read from reg A
		c.RegA = c.RegA >> operand
	case 1: // The bxl instruction (opcode 1) XOR of reg B and literal operand, store in reg B
		c.RegB = c.RegB ^ uint64(x)
	case 2: // The bst instruction (opcode 2) mod 8 combo operand, store in reg B
		c.RegB = operand % 8
	case 3: // The jnz instruction (opcode 3) if reg A == 0, do nothing, else jumps instruction to literal operand
		if c.RegA != 0 {
			c.CurrPtr = uint64(x - 2)
			return
		}
	case 4: // The bxc instruction (opcode 4) XOR of reg B and reg C, store in reg B
		c.RegB = c.RegB ^ c.RegC
	case 5: // ! The out instruction (opcode 5) mod 8 of combo operand
		c.OutInstr = append(c.OutInstr, int(operand%8))
	case 6: // The bdv instruction (opcode 6) like (opcode 0) store in reg B, numerator read from reg A
		c.RegB = c.RegA >> operand
	case 7: // The cdv instruction (opcode 7) like (opcode 0) store in reg C, numerator read from reg A
		c.RegC = c.RegA >> operand
	}
}

func main() {
	computer := ParsePuzzleInput(false, "day17.txt")
	fmt.Printf("Part 1: %s\n", Part1(computer))
	computer = ParsePuzzleInput(false, "day17.txt")
	fmt.Printf("Part 2: %d\n", Part2(computer))
}

func ParsePuzzleInput(sample bool, filename string) *Computer {
	// function to parse the puzzle input from file
	file := utils.GetPuzzleInput(filename, sample)

	input, _ := io.ReadAll(file)
	ptrn := regexp.MustCompile(`(\d+)`)
	matches := ptrn.FindAll(input, -1)
	// Convert matches to integers
	var numbers []int
	for _, match := range matches {
		num, err := strconv.Atoi(string(match)) // Convert byte slice to string, then to int
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			continue
		}
		numbers = append(numbers, num)
	}

	if len(numbers) < 3 {
		panic("parsed input incorrectly")
	}

	return &Computer{
		RegA:     uint64(numbers[0]),
		RegB:     uint64(numbers[1]),
		RegC:     uint64(numbers[2]),
		ProInstr: numbers[3:],
		OutInstr: []int{},
	}
}

func Part1(c *Computer) string {
	c.execProgram()
	resultOut := make([]string, len(c.OutInstr))
	for i, num := range c.OutInstr {
		resultOut[i] = strconv.Itoa(num)
	}

	return strings.Join(resultOut, ",")
}

func Part2(c *Computer) uint64 {
	reset := func(comp *Computer, a uint64) {
		comp.RegA = a
		comp.RegB = 0
		comp.RegC = 0
		comp.OutInstr = []int{}
	}
	max := func(a, b int) int {
		if a < b {
			return b
		}

		return a
	}

	plen := len(c.ProInstr)
	regA := uint64(math.Pow(8, float64(plen-1)))
	power := plen - 2
	match := c.ProInstr[plen-1:]
	output := []int{}
	fmt.Println("PART 2: GO!")
	for !reflect.DeepEqual(output, match) {
		regA += uint64(math.Pow(8, float64(power)))
		reset(c, regA)
		output = c.execProgram()
		fmt.Printf("RegA: %-19d || Power: %-2d || Out: %+v || Next Match: %+v\n", regA, power, output, match)
		if reflect.DeepEqual(output[plen-len(match):], match) {
			power = max(0, power-1)
			match = c.ProInstr[plen-(len(match)+1):]
			fmt.Printf("MATCH FOUND || RegA: %-19d || Power: %-2d || Next Match: %+v\n", regA, power, match)
		}
	}

	return regA
}
