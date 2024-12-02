package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	listA, listB := ParsePuzzleInput()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(listA, listB)
	}
}

func BenchmarkPart2(b *testing.B) {
	listA, listB := ParsePuzzleInput()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Part2(listA, listB)
	}
}
