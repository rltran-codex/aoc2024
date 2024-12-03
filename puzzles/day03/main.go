package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func main() {
	in := utils.GetFlatPuzzleInput("day3.txt", false)
	fmt.Printf("%d\n", Part1(in))
	fmt.Printf("%d\n", Part2(in))
}

func Part1(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0
	match := r.FindAllStringSubmatch(input, -1)

	for _, v := range match {
		x, _ := strconv.Atoi(v[1])
		y, _ := strconv.Atoi(v[2])
		result += (x * y)
	}

	return result
}

func Part2(input string) int {
	r1 := regexp.MustCompile(`do\(\)`)
	r2 := regexp.MustCompile(`don't\(\)`)
	r3 := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0

	cmd := mapDoAndDont(r1.FindAllStringIndex(input, -1), r2.FindAllStringIndex(input, -1))
	n1 := r3.FindAllStringSubmatchIndex(input, -1)
	n := r3.FindAllStringSubmatch(input, -1)

	calculate := true
	// anything before the first "dont", multiply
	// anything between "dont" - "do", do not multiply

	// current multiplier index
	midx := 0
	sidx := len(n)
	for i := range input {
		// stop if no more product operands
		if midx == sidx {
			break
		}
		if c, ok := cmd[i]; ok {
			calculate = c // update command
		}

		if i == n1[midx][0] {
			if calculate {
				x, _ := strconv.Atoi(n[midx][1])
				y, _ := strconv.Atoi(n[midx][2])
				result += (x * y)
			}
			midx += 1
		}
	}
	return result
}

func mapDoAndDont(do [][]int, dont [][]int) map[int]bool {
	m := make(map[int]bool)
	for _, v := range do {
		m[v[0]] = true
	}
	for _, v := range dont {
		m[v[0]] = false
	}

	return m
}
