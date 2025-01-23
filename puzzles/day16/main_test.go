package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	m := ParsePuzzleInput(false, "day16.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(m) // code to benchmark
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	m := ParsePuzzleInput(false, "day16.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(m) // code to benchmark
	}
}

func TestSmallSample(t *testing.T) {
	d := ParsePuzzleInput(true, "day16_small.txt")

	e := 7036
	a := d.findShortestPath()
	if e != a {
		t.Errorf("Dijkstra's Algorithm returned %d, when expecting %d", a, e)
	}
}

// func TestSmallSample2(t *testing.T) {
// 	d := ParsePuzzleInput(true, "day16_small.txt")

// 	e := 45
// 	a := d.findBestSpots()
// 	fmt.Print(e, a)
// }

func TestSample(t *testing.T) {
	d := ParsePuzzleInput(true, "day16.txt")

	e := 11048
	a := d.findShortestPath()
	if e != a {
		t.Errorf("Dijkstra's Algorithm returned %d, when expecting %d", a, e)
	}
}

func TestPaths(t *testing.T) {
	// d := ParsePuzzleInput(true, "day16_small.txt")
}
