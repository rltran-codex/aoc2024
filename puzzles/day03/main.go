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

func product(v1 string, v2 string) int {
	x, _ := strconv.Atoi(v1)
	y, _ := strconv.Atoi(v2)
	return (x * y)
}

func Part1(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	result := 0
	match := r.FindAllStringSubmatch(input, -1)

	for _, v := range match {
		result += product(v[1], v[2])
	}

	return result
}

func Part2(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	result := 0
	matches := r.FindAllStringSubmatch(input, -1)
	calc := true

	for _, v := range matches {
		switch v[0] {
		case "don't()":
			calc = false
		case "do()":
			calc = true
		default:
			if calc {
				result += product(v[1], v[2])
			}
		}
	}

	return result
}
