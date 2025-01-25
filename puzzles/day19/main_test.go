package main

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	avail, design := ParsePuzzleInput(false, "day19.txt")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// code to benchmark
		Part1(avail, design)
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
	avail, design := ParsePuzzleInput(true, "day19.txt")
	expected_avail := []string{
		"r", "wr", "b", "g", "bwu", "rb", "gb", "br",
	}
	slices.Sort(expected_avail)
	sort.Slice(expected_avail, func(i, j int) bool {
		return len(expected_avail[i]) > len(expected_avail[j])
	})
	expected_design := []string{
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}

	if !reflect.DeepEqual(expected_avail, avail) {
		t.Errorf("Expected: %+v. Actual: %+v", expected_avail, avail)
	}
	if !reflect.DeepEqual(expected_design, design) {
		t.Errorf("Expected: %+v. Actual: %+v", expected_design, design)
	}
}

func TestPart1(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 6
	a := Part1(avail, design)
	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCase1(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a, _ := matchDesign(avail, design[0])

	if a != e {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase2(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a, _ := matchDesign(avail, design[1])

	if a != e {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase3(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a, _ := matchDesign(avail, design[2])

	if a != e {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase4(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a, _ := matchDesign(avail, design[3])

	if a != 1 {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase5(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	a, _ := matchDesign(avail, design[4])

	if a != 0 {
		t.Errorf("Expected %+v. Actual %+v", nil, a)
	}
}

func TestCase6(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a, _ := matchDesign(avail, design[5])

	if a != e {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase7(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	a, _ := matchDesign(avail, design[6])

	if a != e {
		t.Errorf("Expected %+v. Actual %+v", e, a)
	}
}

func TestCase8(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	a, _ := matchDesign(avail, design[7])

	if a != 0 {
		t.Errorf("Expected %+v. Actual %+v", nil, a)
	}
}

func TestCaseAll1(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 2
	_, a := matchDesign(avail, design[0])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll2(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	_, a := matchDesign(avail, design[1])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll3(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 4
	_, a := matchDesign(avail, design[2])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll4(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 6
	_, a := matchDesign(avail, design[3])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll5(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 0
	_, a := matchDesign(avail, design[4])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll6(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 1
	_, a := matchDesign(avail, design[5])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll7(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 2
	_, a := matchDesign(avail, design[6])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestCaseAll8(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 0
	_, a := matchDesign(avail, design[7])

	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}

func TestPart2Sample(t *testing.T) {
	avail, design := ParsePuzzleInput(true, "day19.txt")
	e := 16
	a := Part2(avail, design)
	if a != e {
		t.Errorf("Expected: %d. Actual: %d.", e, a)
	}
}
