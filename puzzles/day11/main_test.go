package main

import (
	"fmt"
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
		Stones:     make(map[string][]string),
		StoneCount: len(stones),
	}
	ps.Stones["0"] = append(ps.Stones["0"], "1")

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
	ps := PlutoStones{
		Queue:      stones,
		Stones:     make(map[string][]string),
		StoneCount: len(stones),
	}
	ps.Stones["0"] = append(ps.Stones["0"], "1")

	expected := []int{3, 4, 5, 9, 13, 22}
	expectedArr := [][]string{
		{"253000", "1", "7"},
		{"253", "0", "2024", "14168"},
		{"512072", "1", "20", "24", "28676032"},
		{"512", "72", "2024", "2", "0", "2", "4", "2867", "6032"},
		{"1036288", "7", "2", "20", "24", "4048", "1", "4048", "8096", "28", "67", "60", "32"},
		{"2097446912", "14168", "4048", "2", "0", "2", "4", "40", "48", "2024", "40", "48", "80", "96", "2", "8", "6", "7", "6", "0", "3", "2"},
	}
	for i := 0; i < 6; i++ {
		ps.blink(1)
		fmt.Printf("After %d blink: %+v\n", i+1, ps.Queue)
		if ps.StoneCount != expected[i] {
			t.Errorf("Expected %d, but was %d", expected[i], ps.StoneCount)
			for j := 0; j < len(ps.Queue); j++ {
				if expectedArr[i][j] != ps.Queue[j] {
					t.Errorf("Expected %s, but was %s at index %d", expectedArr[i][j], ps.Queue[j], j)
				}
			}
		}
	}
}
