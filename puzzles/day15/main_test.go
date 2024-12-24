package main

import (
	"bufio"
	"strings"
	"testing"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func BenchmarkPart1(b *testing.B) {
	data := ParsePuzzleInput(false, "day15.txt") // set up dataset (aka puzzle data)
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

func TestMovingFish(t *testing.T) {
	ls := ParsePuzzleInput(true, "day15_sample.txt")
	expected := openExpected("day15_expected.txt")
	max := len(ls.Instructions) - 1

	for i := range ls.Instructions {
		ls.performMove()
		actual := ""
		for _, v := range ls.Warehouse {
			actual += strings.Join(v, "")
		}

		if i != max && expected[i] != actual {
			t.Errorf("Expected map does not match with actual.")
		}
	}
}

func TestSmallSample(t *testing.T) {
	ls := ParsePuzzleInput(true, "day15_sample.txt")
	expected := 2028
	for len(ls.Instructions) > 0 {
		ls.performMove()
	}

	actual := ls.gatherGPS()
	if expected != actual {
		t.Errorf("Expected %d, but actual %d.", expected, actual)
	}
}

func TestBigSample(t *testing.T) {
	ls := ParsePuzzleInput(true, "day15.txt")
	expected := 10092
	for len(ls.Instructions) > 0 {
		ls.performMove()
	}

	actual := ls.gatherGPS()
	if expected != actual {
		t.Errorf("Expected %d, but actual %d.", expected, actual)
	}
}

func openExpected(filename string) []string {
	file := utils.GetPuzzleInput(filename, true)
	defer file.Close()

	scn := bufio.NewScanner(file)
	cases := []string{}

	c := ""
	for scn.Scan() {
		line := strings.TrimSpace(scn.Text())
		if len(line) == 0 {
			cases = append(cases, c)
			c = ""
			continue
		}

		c += line
	}

	return cases
}
