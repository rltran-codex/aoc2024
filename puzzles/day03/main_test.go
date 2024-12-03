package main

import (
	"testing"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func BenchmarkPart1(b *testing.B) {
	in := utils.GetFlatPuzzleInput("day3.txt", false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(in)
	}
}

func BenchmarkPart2(b *testing.B) {
	in := utils.GetFlatPuzzleInput("day3.txt", false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(in)
	}
}
