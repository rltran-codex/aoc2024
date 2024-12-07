package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	data, start := ParsePuzzleInput()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part1(data, start)
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}
