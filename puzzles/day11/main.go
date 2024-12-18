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
	StonesMemo map[string]int
	StoneCount int
}

func main() {
	// main area to display puzzle answers
	ps := ParsePuzzleInput()
	fmt.Printf("Part 1: %d\n", Part1(ps))
	ps = ParsePuzzleInput()

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
		StonesMemo: make(map[string]int),
		Queue:      nums,
		StoneCount: len(nums),
	}

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
	ps.StoneCount = 0
	for i := 0; i < len(ps.Queue); i++ {
		currStone := ps.Queue[i]
		ps.StoneCount += ps.recurBlink(0, k, currStone)
	}
}

func (ps *PlutoStones) recurBlink(currBlink int, maxBlink int, stone string) int {
	if currBlink == maxBlink {
		return 1
	}
	currKey := strings.Join([]string{stone, strconv.Itoa(currBlink)}, ";")
	val, ok := ps.StonesMemo[currKey]
	if ok {
		return val
	}

	count := 0
	if stone == "0" {
		count = ps.recurBlink(currBlink+1, maxBlink, "1")
		ps.StonesMemo[currKey] = count
		return count
	}

	if len(stone)%2 == 0 {
		evo := handleEven(stone)
		count = ps.recurBlink(currBlink+1, maxBlink, evo[0]) + ps.recurBlink(currBlink+1, maxBlink, evo[1])
	} else {
		evo := handleMultiply(stone)[0]
		count = ps.recurBlink(currBlink+1, maxBlink, evo)
	}
	ps.StonesMemo[currKey] = count
	return count
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
