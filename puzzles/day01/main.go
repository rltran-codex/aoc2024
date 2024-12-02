package main

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func main() {
	listA, listB := ParsePuzzleInput()

	fmt.Printf("Total distance: %d\n", Part1(listA, listB))
	fmt.Printf("Similarity score: %d\n", Part2(listA, listB))
}

func ParsePuzzleInput() ([]int, []int) {
	// file := utils.GetPuzzleInput("day1.txt", false)
	file := utils.GetPuzzleInput("day1.txt", false)
	defer file.Close()

	r := regexp.MustCompile(`^(\d+)\s*(\d+)$`)
	var listA []int
	var listB []int
	scn := bufio.NewScanner(file)

	for scn.Scan() {
		line := scn.Text()
		match := r.FindAllStringSubmatch(line, -1)

		if match != nil {
			num1, _ := strconv.Atoi(match[0][1])
			num2, _ := strconv.Atoi(match[0][2])
			listA = append(listA, num1)
			listB = append(listB, num2)
		}
	}

	return listA, listB
}

func Part1(listA []int, listB []int) int {
	sort.Ints(listA)
	sort.Ints(listB)
	result := 0

	for i := 0; i < len(listA); i++ {
		dist := math.Abs(float64(listA[i]) - float64(listB[i]))
		result += int(dist)
	}

	return result
}

func Part2(listA []int, listB []int) int {
	amap := make(map[int]int)
	bmap := make(map[int]int)
	result := 0

	// before calculating score, map the freq to each number
	for i := 0; i < len(listB); i++ {
		_, ok := bmap[listB[i]]
		if !ok {
			bmap[listB[i]] = 1
		} else {
			bmap[listB[i]] += 1
		}
	}

	for i := 0; i < len(listA); i++ {
		n, ok := amap[listA[i]]

		// not found in amap
		if !ok {
			freq := bmap[listA[i]]
			amap[listA[i]] = freq

			n = amap[listA[i]]
		}

		result += (listA[i] * n)
	}

	return result
}
