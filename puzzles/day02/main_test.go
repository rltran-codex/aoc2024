package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	listA := ParsePuzzleInput()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(listA)
	}
}

func BenchmarkPart2(b *testing.B) {
	listA := ParsePuzzleInput()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(listA)
	}
}
