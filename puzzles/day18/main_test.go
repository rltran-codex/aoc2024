package main

import (
	"strings"
	"testing"
)

var GridSize int = 6

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	input := ParsePuzzleInput(false, "day18.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(70, input, 1026) // code to benchmark
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	input := ParsePuzzleInput(false, "day18.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(70, input, 1026) // code to benchmark
	}
}

func TestGenerateGrid(t *testing.T) {
	bytes := ParsePuzzleInput(true, "day18.txt")
	grid := generateGrid(GridSize, bytes[:12])
	expected := []string{
		"...#...",
		"..#..#.",
		"....#..",
		"...#..#",
		"..#..#.",
		".#..#..",
		"#.#....",
	}

	for i, v := range grid {
		row := strings.Join(v, "")
		if expected[i] != row {
			t.Errorf("Expected '%s', but was '%s' for %d row", expected[i], row, i)
		}
	}
}

func TestPathfinder(t *testing.T) {
	bytes := ParsePuzzleInput(true, "day18.txt")
	grid := generateGrid(GridSize, bytes[:12])
	expected := []string{
		"OO.#OOO",
		".O#OO#O",
		".OOO#OO",
		"...#OO#",
		"..#OO#.",
		".#.O#..",
		"#.#OOOO",
	}

	endPt := Coordinate{
		X: GridSize,
		Y: GridSize,
	}
	startPt := Coordinate{
		X: 0,
		Y: 0,
	}

	path, steps := pathfinder(grid, startPt, endPt)
	if steps != 22 {
		t.Errorf("invalid shortest path. Expected: %d, Actual: %d", 22, steps)
	}

	for _, v := range path {
		grid[v.Y][v.X] = "O"
	}
	for i, v := range grid {
		row := strings.Join(v, "")
		if expected[i] != row {
			t.Errorf("Expected '%s', but was '%s' for %d row", expected[i], row, i)
		}
	}
}

func TestPart2(t *testing.T) {
	bytes := ParsePuzzleInput(true, "day18.txt")
	e := "6,1"
	a := Part2(GridSize, bytes, 12)
	if a != e {
		t.Errorf("Expected '%s'. Actual '%s'.", e, a)
	}
}
