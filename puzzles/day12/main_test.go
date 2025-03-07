package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	// data := ParsePuzzleInput(false, "day12.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		// Part1(data)
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	data := ParsePuzzleInput(false, "day12.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part2(data)
	}
}

func TestFence(t *testing.T) {
	garden := ParsePuzzleInput(true, "day12.txt")
	n, m := len(garden), len(garden[0])
	visited := make([][]bool, n)
	for i := range garden {
		visited[i] = make([]bool, m)
	}

	area, perimeter, _ := fence(garden, visited, "R", [2]int{0, 0})
	if area != 12 {
		t.Error("expected area = 12 but was", area)
	}
	if perimeter != 18 {
		t.Error("expected perimeter = 18 but was", perimeter)
	}
}

func TestCornerCalc(t *testing.T) {
	garden := ParsePuzzleInput(true, "day12_part2_1.txt")
	n, m := len(garden), len(garden[0])
	visited := make([][]bool, n)
	for i := range garden {
		visited[i] = make([]bool, m)
	}

	_, _, c := fence(garden, visited, "C", [2]int{2, 1})
	if c != 8 {
		t.Error("expected 8 sides, but was ", c)
	}
}

func TestSample1(t *testing.T) {
	garden := ParsePuzzleInput(true, "day12_part2_1.txt")
	cost := Part2(garden)
	if cost != 80 {
		t.Error("80 != ", cost)
	}
}
