package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Lanternfish struct {
	grid         [][]string
	pos          *[2]int
	instructions []string
}

type space struct {
	v string
	c [2]int
}

func main() {
	// main area to display puzzle answers
	grid, instr, fishPos := ParsePuzzleInput(false, "day15.txt")
	fish := Lanternfish{
		grid:         grid,
		instructions: instr,
		pos:          &fishPos,
	}
	fmt.Printf("Part 1: %d\n", Part1(fish))
	// main area to display puzzle answers
	grid, instr, fishPos = ParsePuzzleInput(false, "day15.txt")
	fish = Lanternfish{
		grid:         grid,
		instructions: instr,
		pos:          &fishPos,
	}
	fish.expandGrid()
	fmt.Printf("Part 2: %d\n", Part2())
}

func ParsePuzzleInput(sample bool, filename string) ([][]string, []string, [2]int) {
	file := utils.GetPuzzleInput(filename, sample)
	scn := bufio.NewScanner(file)
	var grid [][]string
	var instr []string
	var curr [2]int

	moveInstructions := false
	for scn.Scan() {
		line := scn.Text()
		if len(strings.TrimSpace(line)) == 0 {
			moveInstructions = true
		}

		l := strings.Split(line, "")
		if moveInstructions {
			instr = append(instr, l...)
		} else {
			grid = append(grid, l)
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "@" {
				curr = [2]int{i, j}
			}
		}
	}
	return grid, instr, curr
}

func Part1(f Lanternfish) int {
	for i := range f.instructions {
		instr := f.instructions[i]
		f.moveFish(instr)
	}

	sum := 0
	for r := range f.grid {
		for c := range f.grid[r] {
			if f.grid[r][c] == "O" {
				sum += 100*r + c
			}
		}
	}

	return sum
}

func Part2() int {
	panic("Not yet implemented")
}

func parseMove(m string) (dy int, dx int) {
	switch m {
	case "<":
		dy = 0
		dx = -1
	case ">":
		dy = 0
		dx = 1
	case "^":
		dy = -1
		dx = 0
	case "v":
		dy = 1
		dx = 0
	}
	return dy, dx
}

func (fish *Lanternfish) performPush(stack []space) {
	for len(stack) > 2 {
		e1 := utils.PopQueue(&stack)
		e2 := stack[0]

		y, x := e1.c[0], e1.c[1]
		fish.grid[y][x] = e2.v
	}

	e1 := utils.PopQueue(&stack)
	e2 := stack[0]
	y, x := e1.c[0], e1.c[1]
	fish.grid[y][x] = e2.v
	if e2.v == "@" {
		fish.pos = &[2]int{y, x}
	}

	y, x = e2.c[0], e2.c[1]
	fish.grid[y][x] = "."
}

func (fish *Lanternfish) moveFish(instruction string) {
	y, x := fish.pos[0], fish.pos[1]
	dy, dx := parseMove(instruction)

	// use a stack
	stack := []space{}
	for {
		currTile := fish.grid[y][x]
		if currTile == "#" {
			return
		}

		utils.PushStack(&stack, space{
			v: currTile,
			c: [2]int{y, x},
		})

		if currTile == "." {
			fish.performPush(stack)
			break
		}
		y, x = (y + dy), (x + dx)
	}
}

func (fish *Lanternfish) expandGrid() {
	expandedGrid := [][]string{}
	for r := range fish.grid {
		row := []string{}
		for c := range fish.grid[r] {
			var t []string
			switch fish.grid[r][c] {
			case "O":
				t = []string{"[", "]"}
			case ".":
				t = []string{".", "."}
			case "@":
				t = []string{"@", "."}
				// update fish.pos
				fish.pos = &[2]int{r, (2 * c)}

			case "#":
				t = []string{"#", "#"}
			}
			row = append(row, t...)
		}
		expandedGrid = append(expandedGrid, row)
	}

	fish.grid = expandedGrid
}

func (fish *Lanternfish) moveFish2(tile [2]int, instruction string) {
	x, y := tile[0], tile[1]
	dy, dx := parseMove(instruction)

	// use a stack
	stack := []space{}
	for {
		currTile := fish.grid[y][x]
		if currTile == "#" {
			return
		}

		utils.PushStack(&stack, space{
			v: currTile,
			c: [2]int{y, x},
		})

		switch currTile {
		case "[", "]":
			if instruction == "<" || instruction == ">" {
				return
			}

		case "O":
			continue
		case ".":
			// space available, perform push
		}
		y, x = (y + dy), (x + dx)
	}
}
