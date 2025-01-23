package main

import (
	"fmt"
	"reflect"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	// set up dataset (aka puzzle data)
	data := ParsePuzzleInput(false, "day17.txt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		data.execProgram()
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
	comp := ParsePuzzleInput(true, "day17.txt")
	print(comp)
	if comp.RegA != 729 {
		t.Errorf("Expected 729, but Reg A was %d", comp.RegA)
	}
	if comp.RegB != 0 {
		t.Errorf("Expected 0, but Reg B was %d", comp.RegB)
	}
	if comp.RegC != 0 {
		t.Errorf("Expected 0, but Reg C was %d", comp.RegC)
	}

	expected := []int{0, 1, 5, 4, 3, 0}
	for i := range comp.ProInstr {
		if comp.ProInstr[i] != expected[i] {
			t.Errorf("Expected instruction at idx %d to be %d, but was %d", i, expected[i], comp.ProInstr[i])
		}
	}
}

func TestCase1(t *testing.T) {
	computer := Computer{
		RegA: 0,
		RegB: 0,
		RegC: 9,
	}

	computer.handleOpcode(2, 6)
	if computer.RegB != 1 {
		t.Errorf("Expected RegB to be 1, but is %d", computer.RegB)
	}

}

func TestCase2(t *testing.T) {
	computer := Computer{
		RegA:     10,
		RegB:     0,
		RegC:     0,
		ProInstr: []int{5, 0, 5, 1, 5, 4},
		OutInstr: []int{},
	}

	computer.execProgram()
	expect := []int{0, 1, 2}
	if !reflect.DeepEqual(computer.OutInstr, expect) {
		t.Errorf("Expected %+v, but was %+v", expect, computer.OutInstr)
	}
}

func TestCase3(t *testing.T) {
	computer := Computer{
		RegA:     2024,
		RegB:     0,
		RegC:     0,
		ProInstr: []int{0, 1, 5, 4, 3, 0},
		OutInstr: []int{},
	}
	computer.execProgram()
	expectedOut := []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}
	expectedA := 0
	if !reflect.DeepEqual(expectedOut, computer.OutInstr) {
		t.Errorf("Expected %+v, but was %+v", expectedOut, computer.OutInstr)
	}

	if computer.RegA != expectedA {
		t.Errorf("Expected RegA to be 0, but is %d", computer.RegA)
	}
}

func TestCase4(t *testing.T) {
	computer := Computer{
		RegA: 0,
		RegB: 29,
		RegC: 0,
	}

	computer.handleOpcode(1, 7)
	if computer.RegB != 26 {
		t.Errorf("Expected RegB to be 26, but is %d", computer.RegB)
	}
}

func TestCase5(t *testing.T) {
	computer := Computer{
		RegA:     0,
		RegB:     2024,
		RegC:     43690,
		ProInstr: []int{4, 0},
		OutInstr: []int{},
	}

	computer.execProgram()
	if computer.RegB != 44354 {
		t.Errorf("Expected RegB to be 44354, but was %d", computer.RegB)
	}
}

func TestCasePart2Sample(t *testing.T) {
	computer := Computer{
		RegA:     2024,
		RegB:     0,
		RegC:     0,
		ProInstr: []int{0, 3, 5, 4, 3, 0},
		OutInstr: []int{},
	}
	fmt.Printf("%+v\n", computer)
	regA := Part2(&computer)
	if regA != 117440 {
		t.Errorf("invalid: %d", regA)
	}
	fmt.Printf("regA: %d, %+v", regA, computer)
}
