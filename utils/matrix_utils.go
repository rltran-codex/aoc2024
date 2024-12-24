package utils

import (
	"bufio"
	"strings"
)

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
