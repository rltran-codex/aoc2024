package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func main() {
	data := ParsePuzzleInput()
	fmt.Printf("Safe reports 1: %d\n", Part1(data))
	fmt.Printf("Safe reports 2: %d\n", Part2(data))
}

func ParsePuzzleInput() [][]int {
	file := utils.GetPuzzleInput("day2.txt", false)
	defer file.Close()
	var reports [][]int

	scn := bufio.NewScanner(file)
	for scn.Scan() {
		line := strings.TrimSpace(scn.Text())

		arr := strings.Split(line, " ")
		nums := make([]int, len(arr))
		for i := 0; i < len(arr); i++ {
			num, _ := strconv.Atoi(arr[i])
			nums[i] = num
		}
		reports = append(reports, nums)
	}

	return reports
}

func Part1(reports [][]int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	results := 0

	// use waitgroup to quickly iterate each row in 2D array.
	for _, v := range reports {
		wg.Add(1)
		go func(nums []int) {
			defer wg.Done()

			n := recurValidate(nums, 0, nums[1]-nums[0] > 0)
			if n == -1 {
				// successfully passes conditions, increment results
				mu.Lock()
				results += 1
				mu.Unlock()
			}
		}(v)
	}

	wg.Wait()
	return results
}

func Part2(reports [][]int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	results := 0

	for _, nums := range reports {
		wg.Add(1)
		go func(v []int) {
			defer wg.Done()

			n := recurValidate(v, 0, v[1]-v[0] > 0)

			// first recursion was successful
			if n == -1 {
				mu.Lock()
				results += 1
				mu.Unlock()
				return
			}

			// unsafe report, try and permutate around area of failure
			delIdx := []int{n - 1, n, n + 1}
			for _, d := range delIdx {
				if d < 0 || d >= len(v) {
					continue
				}

				t1 := utils.RemoveIndex(v, d)
				n = recurValidate(t1, 0, t1[1]-t1[0] > 0)
				if n == -1 {
					mu.Lock()
					results += 1
					mu.Unlock()
					return
				}
			}
		}(nums)
	}

	wg.Wait()

	return results
}

/*
	  Function validates report, if unsafe the index of where it fails is returned, else -1.
		  If lvl diff violates any of the two conditions:
			The levels are either all increasing or all decreasing.
			Any two adjacent levels differ by at least one and at most three.
*/
func recurValidate(nums []int, idx int, isIncreasing bool) int {
	// base case 1: if reach the end of the array
	if idx >= (len(nums) - 1) {
		return -1
	}

	lvlChange := float64(nums[idx+1] - nums[idx])
	diff := math.Abs(lvlChange)

	if diff > 3 || diff < 1 {
		return idx
	}

	if lvlChange > 0 && !isIncreasing {
		return idx
	} else if lvlChange < 0 && isIncreasing {
		return idx
	}

	return recurValidate(nums, idx+1, isIncreasing)
}
