package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func GetPuzzleInput(fn string, useSample bool) *os.File {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: Unable to determine the current file")
	}

	currDir := filepath.Dir(filename)
	var inDir string
	if useSample {
		inDir = filepath.Join(currDir, "input", "sample", fn)
	} else {
		inDir = filepath.Join(currDir, "input", fn)
	}

	file, err := os.Open(inDir)
	if err != nil {
		panic(err)
	}

	return file
}

func GetFlatPuzzleInput(fn string, useSample bool) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: Unable to determine the current file")
	}

	currDir := filepath.Dir(filename)
	var inDir string
	if useSample {
		inDir = filepath.Join(currDir, "input", "sample", fn)
	} else {
		inDir = filepath.Join(currDir, "input", fn)
	}

	data, err := os.ReadFile(inDir)
	if err != nil {
		panic(err)
	}

	input := strings.ReplaceAll(string(data), "\n", "")
	return input
}

func RemoveIndex(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}
