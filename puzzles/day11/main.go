package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type PlutoStones struct {
	Queue      []string
	Stones     map[string][]string
	StoneCount int
}

func main() {
	// main area to display puzzle answers
	ps := ParsePuzzleInput()
	fmt.Printf("Part 1: %d\n", Part1(ps))
	fmt.Printf("Part 2: %d\n", Part2(ps))
}

func ParsePuzzleInput() *PlutoStones {
	// function to parse the puzzle input from file
	file := utils.GetPuzzleInput("day11.txt", false)
	defer file.Close()

	scn := bufio.NewScanner(file)
	scn.Scan()
	nums := strings.Split(scn.Text(), " ")
	ps := PlutoStones{
		Stones:     make(map[string][]string),
		Queue:      nums,
		StoneCount: len(nums),
	}
	ps.Stones["0"] = append(ps.Stones["0"], "1")
	return &ps
}

func Part1(ps *PlutoStones) int {
	ps.blink(25)

	return ps.StoneCount
}

func Part2(ps *PlutoStones) int {
	ps.blink(75)

	return ps.StoneCount
}

func (ps *PlutoStones) blink(k int) {
	for i := 0; i < k; i++ {
		sz := len(ps.Queue)

		for j := 0; j < sz; j++ {
			// pop the first element
			currStone := ps.Queue[0]
			ps.Queue = ps.Queue[1:]
			// check if it exists in map
			evoStones, ok := ps.Stones[currStone]

			// handle the evolution cycle of the stone
			if !ok {
				if len(currStone)%2 == 0 {
					evoStones = handleEven(currStone)
				} else {
					evoStones = handleMultiply(currStone)
				}
			}

			// increment if a new stone was added
			if len(evoStones) > 1 {
				ps.StoneCount++
			}
			ps.Queue = append(ps.Queue, evoStones...)
			ps.Stones[currStone] = evoStones
		}
	}
}

func handleEven(stone string) []string {
	mid := len(stone) / 2
	left := strings.TrimLeft(stone[:mid], "0")
	right := strings.TrimLeft(stone[mid:], "0")

	if len(left) == 0 {
		left = "0"
	}
	if len(right) == 0 {
		right = "0"
	}
	return []string{left, right}
}

func handleMultiply(stone string) []string {
	n, _ := strconv.Atoi(stone)
	n *= 2024
	return []string{strconv.Itoa(n)}
}
