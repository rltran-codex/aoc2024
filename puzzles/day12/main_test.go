package main

import (
	"fmt"
	"testing"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	data := ParsePuzzleInput(false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part1(data)
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}

func TestArea(t *testing.T) {
	garden := ParsePuzzleInput(true)
	expected := []string{
		"R: 12",
		"I: 4",
		"C: 14",
		"F: 10",
		"V: 13",
		"J: 11",
		"C: 1",
		"E: 13",
		"I: 14",
		"M: 5",
		"S: 3",
	}
	for _, v := range garden.FarmGraph.Nodes {
		currflower := v.Value
		area, _, err := garden.identifyRegion(v)
		if err != nil {
			continue
		}
		result := fmt.Sprintf("%s: %d", currflower, area)
		if utils.Index(expected, result) == -1 {
			t.Errorf("Could not locate expected value for expected region %s", result)
		} else {
			t.Logf("Area traversal with node at [%d, %d] was successful: %s", v.X, v.Y, result)
		}
	}
}

func TestPerimeter(t *testing.T) {
	garden := ParsePuzzleInput(true)
	expected := []string{
		"R: 18",
		"I: 8",
		"C: 28",
		"F: 18",
		"V: 20",
		"J: 20",
		"C: 4",
		"E: 18",
		"I: 22",
		"M: 12",
		"S: 8",
	}
	for _, v := range garden.FarmGraph.Nodes {
		currflower := v.Value
		_, perimeter, err := garden.identifyRegion(v)
		if err != nil {
			continue
		}
		result := fmt.Sprintf("%s: %d", currflower, perimeter)
		if utils.Index(expected, result) == -1 {
			t.Errorf("Could not locate expected value for expected region %s", result)
		} else {
			t.Logf("Perimeter traversal with node at [%d, %d] was successful: %s", v.X, v.Y, result)
		}
	}
}

func TestSides(t *testing.T) {
	g := ParsePuzzleInput(true)

	// expected := []string{
	// 	"A: 4",
	// 	"B: 4",
	// 	"C: 8",
	// 	"D: 4",
	// 	"E: 4",
	// }

	for _, v := range g.FarmGraph.Nodes {
		g.identifyRegion(v)
		fmt.Print(g)
	}
}

func TestSample1(t *testing.T) {
	g := ParsePuzzleInput(true)
	expected := 1930
	actual := 0
	for _, n := range g.FarmGraph.Nodes {
		area, perimeter, _ := g.identifyRegion(n)
		actual += area * perimeter
	}

	if actual != expected {
		t.Errorf("Expected sample to be %d, but was %d.", expected, actual)
	}
}
