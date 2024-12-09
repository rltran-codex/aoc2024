package utils

import (
	"bufio"
	"os"
	"path/filepath"
	"runtime"
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

func Index[T comparable](slice []T, target T) int {
	for i, item := range slice {
		if item == target {
			return i
		}
	}
	return -1
}

func Swap(slice []string, first int, second int) {
	slice[first], slice[second] = slice[second], slice[first]
}

func GetFlatPuzzleInput(fn string, useSample bool) string {
	data, err := os.ReadFile(buildFilePath(fn, useSample))
	if err != nil {
		panic(err)
	}

	input := strings.ReplaceAll(string(data), "\n", "")
	return input
}

// Opens the puzzle input file and returns 2D array.
func Get2DPuzzleInput(fn string, useSample bool) [][]string {
	lines := GetPuzzleInput(fn, useSample)

	scn := bufio.NewScanner(lines)
	var data [][]string
	for scn.Scan() {
		data = append(data, strings.Split(scn.Text(), ""))
	}

	return data
}

func RemoveIndex[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	ret := make([]T, 0, len(slice)-1)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

func Rotate2DSlice[T any](slice [][]T, dir RotateDirection) [][]T {
	colSize := len(slice[0])
	rowSize := len(slice)
	nMatrix := make([][]T, colSize)
	for i := range nMatrix {
		nMatrix[i] = make([]T, rowSize)
	}

	switch dir {
	case CLOCKWISE:
		for i := 0; i < rowSize; i++ {
			for j := 0; j < colSize; j++ {
				nMatrix[j][rowSize-1-i] = slice[i][j]
			}
		}
	case ANTICLOCK:
		for i := 0; i < rowSize; i++ {
			for j := 0; j < colSize; j++ {
				nMatrix[colSize-1-j][i] = slice[i][j]
			}
		}
	default:
		panic("Unknown direction to rotate 2D array")
	}

	return nMatrix
}

func PopAndRequeue[T any](slice *[]T) T {
	val := (*slice)[0]
	*slice = append((*slice)[1:], val)

	return val
}

func DeepCopy2DArray[T any](original [][]T) *[][]T {
	n := len(original)
	m := len(original[0])
	duplicate := make([][]T, n)
	data := make([]T, n*m)
	for i := range original {
		start := i * m
		end := start + m
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], original[i])
	}

	return &duplicate
}
