package main

import (
	"fmt"
	"testing"
)

const FILE = "day20.txt"

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}

func TestParsing(t *testing.T) {
	track, start := ParsePuzzleInput(true, FILE)
	print(track)
	fmt.Printf("%+v", start)
}

func TestDFS(t *testing.T) {
	track, start := ParsePuzzleInput(true, FILE)
	n, _ := dfs(nil, start[:], 0, track)
	e := 84
	if e != n {
		t.Errorf("Expected %d. Actual: %d", e, n)
	}

	// given 22 because it takes 20 to get to the tile for valid cheat,
	// then +2 to traverse from starting point after phasing throug wall
	n, _ = dfs([]int{7, 9}, []int{7, 11}, 22, track)
	e = 64
	if e != n {
		t.Errorf("Expected wall cheat at %+v distance to be %d, but was %d", []int{7, 11}, e, n)
	}
}

func TestCycle(t *testing.T) {
	track, start := ParsePuzzleInput(true, FILE)
	n, distMap := dfs([]int{3, 3}, start[:], 7, track)
	print(distMap)
	if n != -1 {
		t.Errorf("failed to detect cycle")
	}
}

func TestTraversal(t *testing.T) {
	track, start := ParsePuzzleInput(true, FILE)
	results := traverse(start[:], track)
	eMap := map[int]int{
		2:  14,
		4:  14,
		6:  2,
		8:  4,
		10: 2,
		12: 3,
		20: 1,
		36: 1,
		38: 1,
		40: 1,
		64: 1,
	}

	for k, v := range results {
		e, ok := eMap[k]
		if !ok {
			t.Errorf("Could not find value for number of cheats that save %d picoseconds", k)
		}
		if e != v {
			t.Errorf("When saving %d picoseconds, expected %d, but was %d", k, e, v)
		}
	}
}
