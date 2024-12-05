package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	pageOrdr, printPage := ParsePuzzleInput()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part1(pageOrdr, printPage)
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	pageOrdr, printPage := ParsePuzzleInput()
	_, fSection := Part1(pageOrdr, printPage)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part2(pageOrdr, fSection)
	}
}
