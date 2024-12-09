package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Calibration struct {
	ExpectResult int
	Nums         []int
}

func main() {
	// main area to display puzzle answers
	data := ParsePuzzleInput()
	fmt.Printf("Part 1: %d\n", Part1(data))
	fmt.Printf("Part 2: %d\n", Part2(data))
}

func Part1(data []Calibration) int {
	results := 0
	for _, c := range data {
		n := deduceEquation(c, false)
		if n != -1 {
			results += c.ExpectResult
		}
	}
	return results
}

func Part2(data []Calibration) int {
	results := 0
	for _, c := range data {
		n := deduceEquation(c, true)
		if n != -1 {
			results += c.ExpectResult
		}
	}
	return results
}

// deduceEquation recursively validates the equation by searching for potential operation positions.
// return -1 if equation not found.
func deduceEquation(c Calibration, concat bool) int {
	if len(c.Nums) == 1 {
		return c.Nums[0]
	}

	// try adding
	val := calculate(c.Nums[0], c.Nums[1], "+")
	n := deduceEquation(Calibration{
		ExpectResult: c.ExpectResult,
		Nums:         append([]int{val}, c.Nums[2:]...),
	}, concat)
	if n == c.ExpectResult {
		return n
	}

	// try multiplying
	val = calculate(c.Nums[0], c.Nums[1], "*")
	n = deduceEquation(Calibration{
		ExpectResult: c.ExpectResult,
		Nums:         append([]int{val}, c.Nums[2:]...),
	}, concat)
	if n == c.ExpectResult {
		return n
	}

	if concat {
		// try concatenating
		val = calculate(c.Nums[0], c.Nums[1], "||")
		n = deduceEquation(Calibration{
			ExpectResult: c.ExpectResult,
			Nums:         append([]int{val}, c.Nums[2:]...),
		}, concat)
		if n == c.ExpectResult {
			return n
		}
	}

	// no equation found
	return -1
}

func calculate(x int, y int, op string) int {
	var v int
	switch op {
	case "+":
		v = x + y
	case "*":
		v = x * y
	case "||":
		nx, ny := strconv.Itoa(x), strconv.Itoa(y)
		nNum := strings.Join([]string{nx, ny}, "")
		v, _ = strconv.Atoi(nNum)
	}

	return v
}

func ParsePuzzleInput() []Calibration {
	file := utils.GetPuzzleInput("day7.txt", false)
	defer file.Close()

	scn := bufio.NewScanner(file)
	var input []Calibration

	for scn.Scan() {
		line := strings.Split(scn.Text(), ":")

		sum, nums := line[0], strings.Split(strings.TrimSpace(line[1]), " ")
		numbers := make([]int, len(nums))
		for i, num := range nums {
			n, _ := strconv.Atoi(num)
			numbers[i] = n
		}

		s, _ := strconv.Atoi(sum)
		input = append(input, Calibration{
			ExpectResult: s,
			Nums:         numbers,
		})
	}

	return input
}
