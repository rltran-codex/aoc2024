package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Direction int

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
	UPRIGHT
	UPLEFT
	DOWNRIGHT
	DOWNLEFT
)

func main() {
	in := ParsePuzzleInput()

	// main area to display puzzle answers
	fmt.Printf("Part 1: %d\n", Part1(in))
	fmt.Printf("Part 2: %d\n", Part2(in))
}

func Part1(in [][]string) int {
	// begin iterating each row to look for X
	var mu sync.Mutex
	var wg sync.WaitGroup
	results := 0
	for row := range in {
		for col := range in {
			if in[row][col] == "X" {
				wg.Add(1)
				go func(puzzle [][]string, row int, col int) {
					defer wg.Done()
					words := []int{
						isXMAS(buildWord(puzzle, row, col, 3, UP)),
						isXMAS(buildWord(puzzle, row, col, 3, DOWN)),
						isXMAS(buildWord(puzzle, row, col, 3, LEFT)),
						isXMAS(buildWord(puzzle, row, col, 3, RIGHT)),
						isXMAS(buildWord(puzzle, row, col, 3, UPRIGHT)),
						isXMAS(buildWord(puzzle, row, col, 3, UPLEFT)),
						isXMAS(buildWord(puzzle, row, col, 3, DOWNRIGHT)),
						isXMAS(buildWord(puzzle, row, col, 3, DOWNLEFT)),
					}

					mu.Lock()
					for _, i := range words {
						results += i
					}
					mu.Unlock()
				}(in, row, col)
			}
		}
	}

	wg.Wait()
	return results
}

func Part2(in [][]string) int {
	// begin iterating each row to look for X
	var mu sync.Mutex
	var wg sync.WaitGroup
	results := 0
	for row := range in {
		for col := range in {
			if in[row][col] == "A" {
				wg.Add(1)
				go func(puzzle [][]string, row int, col int) {
					defer wg.Done()
					// create the cross string, so it can be like MASMAS and etc.
					n := []string{
						strings.TrimLeft(buildWord(puzzle, row, col, 1, UPLEFT), "A"),
						"A",
						strings.TrimLeft(buildWord(puzzle, row, col, 1, DOWNRIGHT), "A"),
						strings.TrimLeft(buildWord(puzzle, row, col, 1, UPRIGHT), "A"),
						"A",
						strings.TrimLeft(buildWord(puzzle, row, col, 1, DOWNLEFT), "A"),
					}
					cross := strings.Join(n, "")
					mu.Lock()
					results += isX_MAS(cross)
					mu.Unlock()
				}(in, row, col)
			}
		}
	}

	wg.Wait()
	return results
}

func ParsePuzzleInput() [][]string {
	// function to parse the puzzle input from file
	file := utils.GetPuzzleInput("day4.txt", false)
	defer file.Close()

	scn := bufio.NewScanner(file)
	var data [][]string
	for scn.Scan() {
		data = append(data, strings.Split(scn.Text(), ""))
	}

	return data
}

func isXMAS(str string) int {
	r := regexp.MustCompile(`XMAS|SAMX`)
	match := r.Find([]byte(str))
	if len(match) != 0 {
		return 1
	} else {
		return 0
	}
}

func isX_MAS(str string) int {
	r := regexp.MustCompile(`MASMAS|MASSAM|SAMSAM|SAMMAS`)
	match := r.Find([]byte(str))
	if len(match) != 0 {
		return 1
	} else {
		return 0
	}
}

// recursive function to build the string omnidirectional
func buildWord(puzzle [][]string, row int, col int, step int, dir Direction) string {
	// base case 1: if row && col > 2D array boundaries
	if (row < 0 || row >= len(puzzle)) || (col < 0 || col >= len(puzzle)) {
		return ""
	}

	// base case 2: if no more steps, return current tile
	if step == 0 {
		return puzzle[row][col]
	}

	switch dir {
	case UP:
		return puzzle[row][col] + buildWord(puzzle, row-1, col, step-1, dir)
	case DOWN:
		return puzzle[row][col] + buildWord(puzzle, row+1, col, step-1, dir)
	case LEFT:
		return puzzle[row][col] + buildWord(puzzle, row, col-1, step-1, dir)
	case RIGHT:
		return puzzle[row][col] + buildWord(puzzle, row, col+1, step-1, dir)
	case UPRIGHT:
		return puzzle[row][col] + buildWord(puzzle, row-1, col+1, step-1, dir)
	case UPLEFT:
		return puzzle[row][col] + buildWord(puzzle, row-1, col-1, step-1, dir)
	case DOWNRIGHT:
		return puzzle[row][col] + buildWord(puzzle, row+1, col+1, step-1, dir)
	case DOWNLEFT:
		return puzzle[row][col] + buildWord(puzzle, row+1, col-1, step-1, dir)
	default:
		panic("Unknown direction, you're crazy.")
	}
}
