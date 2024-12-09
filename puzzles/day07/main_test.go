package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	data := ParsePuzzleInput() // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(data) // code to benchmark
	}
}

func BenchmarkPart2(b *testing.B) {
	data := ParsePuzzleInput() // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(data) // code to benchmark
	}
}
