package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Tile struct {
	Val   string
	Up    *Tile
	Down  *Tile
	Right *Tile
	Left  *Tile
}

func main() {
	// main area to display puzzle answers
	data1, start := ParsePuzzleInput()
	data2, _ := ParsePuzzleInput()

	fmt.Printf("Part 1: %d\n", Part1(data1, start))
	fmt.Printf("Part 2: %d\n", Part2(data2, start))
}

func Part1(data [][]string, start []int) int {
	results := 0
	patrolArea(&data, start, &results)
	return results
}

func patrolArea(area *[][]string, currPos []int, steps *int) {
	// base case 1: if next step is out of bounds, guard is done patrolling
	x, y := currPos[0], currPos[1]
	if x-1 < 0 {
		// count the stepping out of bounds as a step
		*steps++
		return
	}

	// increment distinct steps
	if (*area)[x][y] == "." || (*area)[x][y] == "^" {
		*steps++
	}

	// mark it as stepped through
	(*area)[x][y] = "X"

	// if the next step is a #, rotate map counter clockwise to simulate moving right
	if (*area)[x-1][y] == "#" {
		*area = utils.Rotate2DSlice(*area, utils.ANTICLOCK)
		// handle x and y after matrix transpose
		n := len(*area)
		nx, ny := n-1-y, x
		x, y = nx, ny
	}

	currPos[0] = x - 1
	currPos[1] = y
	// increment steps
	patrolArea(area, currPos, steps)
}

func Part2(data [][]string, start []int) int {
	results := 0
	return results
}

func ParsePuzzleInput() ([][]string, []int) {
	file := utils.GetPuzzleInput("day6.txt", false)
	defer file.Close()

	scn := bufio.NewScanner(file)
	var data [][]string
	startingPosition := make([]int, 2)

	row := 0
	for scn.Scan() {
		line := scn.Text()
		i := strings.Index(line, "^")
		if i != -1 {
			startingPosition[0] = row
			startingPosition[1] = i
		}
		data = append(data, strings.Split(line, ""))

		row++
	}

	return data, startingPosition
}
