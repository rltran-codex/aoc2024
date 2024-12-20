package main

import (
	"fmt"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	d := ParsePuzzleInput(false) // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part1(d) // code to benchmark
	}
}

func BenchmarkPart2(b *testing.B) {
	d := ParsePuzzleInput(false) // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Part2(d) // code to benchmark
	}
}

func TestFindCombo(t *testing.T) {
	test := ClawMachine{
		A: struct {
			X int
			Y int
		}{X: 94, Y: 34},
		B: struct {
			X int
			Y int
		}{X: 22, Y: 67},
		Prize: struct {
			X int
			Y int
		}{X: 8400, Y: 5400},
	}

	actual := test.findCombinations()
	fmt.Println(actual)
}

func TestFindAllCombos(t *testing.T) {
	test := ParsePuzzleInput(true)
	expected := []Combination{
		{N: 80, K: 40},
		{},
		{N: 38, K: 86},
		{},
	}
	for i, v := range test {
		actual := v.findCombinations()
		expect := expected[i]
		if expect.N != actual.N || expect.K != actual.K {
			t.Errorf("FAILED with %v. Expected: %v, but Actual: %v", v, expect, actual)
		}
	}
}

func TestCheapestCost(t *testing.T) {
	test := ParsePuzzleInput(true)
	expected := 480
	actual := 0
	for _, v := range test {
		actual += v.findCheapestCost()
	}

	if expected != actual {
		t.Errorf("Expected total cost is %d, but actual %d", expected, actual)
	}
}

func TestCheapestCostPart2(t *testing.T) {
	test := ParsePuzzleInput(true)
	expected := []bool{
		false,
		true,
		false,
		true,
	}
	for i, v := range test {
		v.Prize.X = 10000000000000 + v.Prize.X
		v.Prize.Y = 10000000000000 + v.Prize.Y
		result := v.findCheapestCost()
		if result > 0 && !expected[i] {
			t.Errorf("Expected total cost is %v, but actual %d", expected[i], result)
		} else if result == 0 && expected[i] {
			t.Errorf("Expected total cost is %v, but actual %d", expected[i], result)
		}
	}
}
