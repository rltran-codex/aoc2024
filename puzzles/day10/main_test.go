package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	_, starting := ParsePuzzleInput() // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(starting) // code to benchmark
	}
}

func BenchmarkPart2(b *testing.B) {
	_, starting := ParsePuzzleInput() // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(starting) // code to benchmark
	}
}
