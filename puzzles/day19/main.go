package main

import (
	"bufio"
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

var cache map[string]int = make(map[string]int)

func matchDesign(availTowels []string, design string) (int, int) {
	var recur func(string) (n int)

	// starting from farthest left
	recur = func(design string) (n int) {
		// check the cache
		if n, ok := cache[design]; ok {
			return n
		}
		defer func() { cache[design] = n }()
		// base case 1: reduced to ""
		if len(design) == 0 {
			return 1
		}

		for i := range availTowels {
			currTowel := availTowels[i]
			c_len := len(currTowel)
			if strings.HasPrefix(design, currTowel) {
				n += recur(design[c_len:])
			}
		}
		return n
	}

	if m := recur(design); m > 0 {
		return 1, m
	} else {
		return 0, 0
	}
}

func main() {
	// main area to display puzzle answers
	availTowels, designs := ParsePuzzleInput(false, "day19.txt")
	fmt.Printf("Part 1: %d\n", Part1(availTowels, designs))
	fmt.Printf("Part 2: %d\n", Part2(availTowels, designs))
}

func ParsePuzzleInput(sample bool, filename string) ([]string, []string) {
	file := utils.GetPuzzleInput(filename, sample)
	scn := bufio.NewScanner(file)

	scn.Scan()
	availTowels := strings.Split(scn.Text(), ", ")
	designs := []string{}
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 0 {
			continue
		}

		designs = append(designs, line)
	}

	// sort alphabetically then by string length descending order
	slices.Sort(availTowels)
	sort.Slice(availTowels, func(a, b int) bool {
		return len(availTowels[a]) > len(availTowels[b])
	})
	return availTowels, designs
}

func Part1(availTowels []string, designs []string) int {
	result := 0
	for _, d := range designs {
		possible, _ := matchDesign(availTowels, d)
		result += possible
	}

	return result
}

func Part2(availTowels []string, designs []string) int {
	combos := 0
	for _, design := range designs {
		_, combo := matchDesign(availTowels, design)
		combos += combo
	}

	return combos
}
