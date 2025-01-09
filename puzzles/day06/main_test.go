package main

import (
	"strings"
	"testing"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	data, start := ParsePuzzleInput(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part1(data, start)
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	data, start := ParsePuzzleInput(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part2(data, start)
	}
}

func TestSamplePart1(t *testing.T) {
	area, start := ParsePuzzleInput(true)
	guard := Guard{
		CurrentPosition:  []int{start[0], start[1]},
		CurrentDirection: DefaultDir,
		Rotation:         DefaultRot,
		PatrolArea:       &area,
	}

	// conduct part 1, get the patrol path
	guard.GetPatrolPath(false)

	// check if path is correct
	expectedMap := utils.Get2DPuzzleInput("day6_path.txt", true)
	for curr := guard.Path; curr != nil; {
		area[curr.X][curr.Y] = "X"
		curr = curr.Next
	}

	for r := 0; r < len(area); r++ {
		for c := 0; c < len(area[0]); c++ {
			if area[r][c] != expectedMap[r][c] {
				t.Errorf("Invalid path point found at [%d, %d]", r, c)
			}
		}
	}

	expected := 41
	actual := guard.DistinctLocations()
	if expected != actual {
		t.Error("Failed sample, did not get 41 distinct postions.")
	}
}

func TestSampleDetection(t *testing.T) {
	area, start := ParsePuzzleInput(true)
	// based on the obstacles from the sample
	expected := map[string]bool{
		"6;3": true,
		"7;6": true,
		"7;7": true,
		"8;1": true,
		"8;3": true,
		"9;7": true,
	}
	actual := make(map[string]bool)

	for o := range expected {
		coor := strings.Split(o, ";")
		x, y := utils.Atoi(coor[0]), utils.Atoi(coor[1])
		g := Guard{
			CurrentPosition:  []int{start[0], start[1]},
			CurrentDirection: DefaultDir,
			Rotation:         DefaultRot,
			PatrolArea:       deepCopy2DArray(area),
			Visited:          make(map[string]bool),
		}

		(*g.PatrolArea)[x][y] = "#"
		if g.GetPatrolPath(true) {
			actual[o] = true
		}
	}

	if len(actual) != len(expected) {
		t.Errorf("Length of actual %d, expected %d", len(actual), len(expected))
	}

	for i := range actual {
		if !expected[i] {
			t.Errorf("Unexpected key found in actual: %s", i)
		}
	}
}

func TestPart2(t *testing.T) {
	area, start := ParsePuzzleInput(true)
	expected := 6
	actual := Part2(area, start)
	if expected != actual {
		t.Errorf("Expected %d, but actual %d", expected, actual)
	}
}
