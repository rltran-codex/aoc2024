package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetPuzzleInput(fn string, isSample bool) *os.File {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: Unable to determine the current file")
	}

	currDir := filepath.Dir(filename)
	var inDir string
	if isSample {
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

func RemoveIndex(slice []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}
