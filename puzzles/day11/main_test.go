package main

import (
	"reflect"
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
	// set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
	}
}

func TestSample1(t *testing.T) {
	stones := []string{
		"0",
		"1",
		"10",
		"99",
		"999",
	}
	ps := PlutoStones{
		Queue:      stones,
		StonesMemo: make(map[string]int),
		StoneCount: len(stones),
	}

	expectedArr := []string{"1", "2024", "1", "0", "9", "9", "2021976"}
	ps.blink(1)
	if !reflect.DeepEqual(expectedArr, ps.Queue) {
		t.Errorf("Expected %+v, but was %+v", expectedArr, ps.Queue)
	}
}

func TestSample2(t *testing.T) {
	stones := []string{
		"125",
		"17",
	}

	expected := []int{3, 4, 5, 9, 13, 22}
	for i := 0; i <= 5; i++ {
		ps := PlutoStones{
			Queue:      stones,
			StonesMemo: make(map[string]int),
			StoneCount: len(stones),
		}
		ps.blink(i + 1)
		if ps.StoneCount != expected[i] {
			t.Errorf("Blink %d: Expected %d, but was %d", i+1, expected[i], ps.StoneCount)
		}
	}
}
