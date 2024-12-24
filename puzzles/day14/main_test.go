package main

import (
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	p := ParsePuzzleInput(false) // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part1(p)
	}
}

func BenchmarkPart2(b *testing.B) {
	// set up dataset (aka puzzle data)
	p := ParsePuzzleInput(false) // set up dataset (aka puzzle data)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part2(p)
	}
}

func TestMovingSample1(t *testing.T) {
	r := Robot{}

	MaxHeight = 7
	MaxWidth = 11
	r.Position.X = 2
	r.Position.Y = 4
	r.Velocity.X = 2
	r.Velocity.Y = -3

	expected := [][]int{
		{4, 1},
		{6, 5},
		{8, 2},
		{10, 6},
		{1, 3},
		{3, 0},
		{5, 4},
		{7, 1},
	}

	for i := 0; i < 7; i++ {
		r.moveBot()
		x := r.Position.X
		y := r.Position.Y
		expect := expected[i]

		if x != expect[0] || y != expect[1] {
			t.Errorf("Expected robot to be at [%d, %d] after %d blink(s), but was found at [%d, %d]", expect[0], expect[1], i+1, x, y)
		}
	}
}

func TestSample(t *testing.T) {
	r := ParsePuzzleInput(true)
	MaxWidth = 11
	MaxHeight = 7
	for _, rbt := range r {
		rbt.moveBotN(100)
	}

	expected := []struct {
		X int
		Y int
	}{
		{X: 0, Y: 2},
		{X: 1, Y: 3},
		{X: 2, Y: 3},
		{X: 6, Y: 0},
		{X: 6, Y: 0},
		{X: 9, Y: 0},
		{X: 5, Y: 4},
		{X: 3, Y: 5},
		{X: 4, Y: 5},
		{X: 4, Y: 5},
		{X: 1, Y: 6},
		{X: 6, Y: 6},
	}

	visited := make(map[*Robot]bool)
	for _, e := range expected {
		for _, rbt := range r {
			if visited[rbt] {
				continue
			}

			if e.X == rbt.Position.X && e.Y == rbt.Position.Y {
				visited[rbt] = true
			}
		}
	}

	if len(visited) != len(r) {
		t.Errorf("Discrepancy found in expected positions vs. actual positions. (%d vs. %d)", len(r), len(visited))
	}

	actual := splitGridAndCount(r)
	expected1 := map[string]int{
		"top_left":  1,
		"top_right": 3,
		"btm_left":  4,
		"btm_right": 1,
	}
	for a := range actual {
		if expected1[a] != actual[a] {
			t.Errorf("Quadrant %s expected %d, actual %d", a, expected1[a], actual[a])
		}
	}
}
