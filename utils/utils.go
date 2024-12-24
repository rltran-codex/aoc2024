package utils

import (
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type RotateDirection int

const (
	CLOCKWISE = iota
	ANTICLOCK
)

func buildFilePath(fn string, sample bool) string {
	_, currFp, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error: Unable to determine the current file")
	}

	currDir := filepath.Dir(currFp)
	var fp string
	if sample {
		fp = filepath.Join(currDir, "input", "sample", fn)
	} else {
		fp = filepath.Join(currDir, "input", fn)
	}

	return fp
}

func GetPuzzleInput(fn string, useSample bool) *os.File {
	file, err := os.Open(buildFilePath(fn, useSample))
	if err != nil {
		panic(err)
	}

	return file
}

func GetFlatPuzzleInput(fn string, useSample bool) string {
	data, err := os.ReadFile(buildFilePath(fn, useSample))
	if err != nil {
		panic(err)
	}

	input := strings.ReplaceAll(string(data), "\n", "")
	return input
}

func Atoi(num string) int {
	n, _ := strconv.Atoi(num)
	return n
}

func ParseKey(x int, y int) string {
	nx := strconv.Itoa(x)
	ny := strconv.Itoa(y)
	return strings.Join([]string{nx, ny}, ";")
}
