package main

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}

func TestA(t *testing.T) {
	grid, instr, fish := ParsePuzzleInput(true, "day15_small.txt")
	fmt.Print(grid, instr, fish)
}

func TestB(t *testing.T) {
	grid, instr, fish := ParsePuzzleInput(true, "day15.txt")
	eFile := utils.GetPuzzleInput("day15_expected.txt", true)
	data, err := io.ReadAll(eFile)
	if err != nil {
		t.Fail()
	}

	replacer := strings.NewReplacer("\r", "", "\n", "")
	e := replacer.Replace(string(data))
	lfish := Lanternfish{
		grid:         grid,
		instructions: instr,
		pos:          &fish,
	}
	for i := range instr {
		m := instr[i]
		lfish.moveFish(m)
	}
	var lines []string
	for _, row := range lfish.grid {
		lines = append(lines, strings.Join(row, ""))
	}
	a := strings.Join(lines, "")
	if a != e {
		t.Error()
	}
}

func TestPart1Small(t *testing.T) {
	grid, instr, fish := ParsePuzzleInput(true, "day15_small.txt")

	lfish := Lanternfish{
		grid:         grid,
		instructions: instr,
		pos:          &fish,
	}
	a := Part1(lfish)
	e := 2028
	if a != e {
		t.Errorf("expevted %d, actual %d", e, a)
	}
}

func TestPart1Large(t *testing.T) {
	grid, instr, fish := ParsePuzzleInput(true, "day15.txt")

	lfish := Lanternfish{
		grid:         grid,
		instructions: instr,
		pos:          &fish,
	}
	a := Part1(lfish)
	e := 10092
	if a != e {
		t.Errorf("expevted %d, actual %d", e, a)
	}
}

func TestExpand(t *testing.T) {
	grid, instr, fish := ParsePuzzleInput(true, "day15.txt")

	lfish := Lanternfish{
		grid:         grid,
		instructions: instr,
		pos:          &fish,
	}
	lfish.expandGrid()
	eY, eX := lfish.pos[0], lfish.pos[1]
	if eY != 4 || eX != 8 {
		t.Error("position not correctly updated after expansion")
	}

}
