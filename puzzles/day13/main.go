package main

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type ClawMachine struct {
	A struct {
		X int
		Y int
	}

	B struct {
		X int
		Y int
	}

	Prize struct {
		X int
		Y int
	}
}

type Combination struct {
	N int
	K int
}

var TokenCostA = 3
var TokenCostB = 1

func main() {
	// main area to display puzzle answers
	data := ParsePuzzleInput(false)
	fmt.Printf("Part 1: %d\n", Part1(data))
	fmt.Printf("Part 2: %d\n", Part2(data))
}

func ParsePuzzleInput(sample bool) []ClawMachine {
	config := utils.GetPuzzleInput("day13.txt", sample)
	defer config.Close()
	scn := bufio.NewScanner(config)

	var input string
	for scn.Scan() {
		input += scn.Text() + "\n"
	}

	clawmachine := []ClawMachine{}
	r := regexp.MustCompile(`Button A[:\s]+X\+(\d+)[,\s]+Y\+(\d+)\s+Button B[:\s]+X\+(\d+)[,\s]+Y\+(\d+)\s+Prize:[:\s]X=(\d+)[,\s]+Y=(\d+)`)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, claw := range matches {
		clawmachine = append(clawmachine, ClawMachine{
			A: struct {
				X int
				Y int
			}{X: Atoi(claw[1]), Y: Atoi(claw[2])},
			B: struct {
				X int
				Y int
			}{X: Atoi(claw[3]), Y: Atoi(claw[4])},
			Prize: struct {
				X int
				Y int
			}{X: Atoi(claw[5]), Y: Atoi(claw[6])},
		})
	}

	return clawmachine
}

func (cm *ClawMachine) findCheapestCost() int {
	c := cm.findCombinations()
	return (c.N * TokenCostA) + (c.K * TokenCostB)
}

func (cm *ClawMachine) findCombinations() Combination {
	// n(ax, ay) + k(bx, by) = (px, py)
	// find the possible n and k using system of equations
	ax, ay := cm.A.X, cm.A.Y
	bx, by := cm.B.X, cm.B.Y
	px, py := cm.Prize.X, cm.Prize.Y
	denominator := ax*by - ay*bx
	if denominator == 0 {
		return Combination{}
	}

	n := float64(px*by-py*bx) / float64(denominator)
	k := (float64(px) - n*float64(ax)) / float64(bx)
	if math.Mod(n, 1) == 0 && math.Mod(k, 1) == 0 {
		return Combination{N: int(n), K: int(k)}
	}
	return Combination{}
}

func Atoi(num string) int {
	n, _ := strconv.Atoi(num)
	return n
}

func Part1(data []ClawMachine) int {
	result := 0
	for _, v := range data {
		result += v.findCheapestCost()
	}

	return result
}
func Part2(data []ClawMachine) int {
	result := 0
	for _, v := range data {
		v.Prize.X = 10000000000000 + v.Prize.X
		v.Prize.Y = 10000000000000 + v.Prize.Y
		result += v.findCheapestCost()
	}

	return result
}
