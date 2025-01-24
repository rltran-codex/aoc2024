package main

import (
	"bufio"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

var memo map[string]bool = make(map[string]bool)

func matchDesign(availTowels []string, design string) map[string]int {
	// start from the far left of the design, try and find a match. if no match found, return nil
	towels := make(map[string]int)

	_, good := recur(availTowels, design, towels)
	// no pattern found
	if good {
		return towels
	} else {
		return nil
	}
}

func findAllCombo(availTowels []string, design string) int {
	towels := make(map[string]int)
	allCombo := []map[string]int{}
	good := recur_allcombo(availTowels, design, towels, &allCombo)

	if good {
		return len(allCombo)
	} else {
		return 0
	}
}

func recur(availTowels []string, design string, combo map[string]int) (map[string]int, bool) {
	// base case 1: reduced to ""
	if len(design) == 0 {
		return combo, true
	}
	for i := range availTowels {
		currTowel := availTowels[i]
		c_len := len(currTowel)
		if c_len > len(design) {
			continue
		}
		ptrn := regexp.MustCompile(fmt.Sprintf("^%s", currTowel))
		match := ptrn.FindAllStringIndex(design[:c_len], 1)
		if match == nil {
			continue
		}

		combo[currTowel] += 1

		// split
		parts := splitUpDesign(match[0], design)
		_, lValid := recur(availTowels, parts[0], combo)
		_, rValid := recur(availTowels, parts[1], combo)

		if lValid && rValid {
			return combo, true
		} else {
			combo[currTowel] -= 1
		}
	}

	return nil, false
}

func recur_allcombo(availTowels []string, design string, combo map[string]int, allCombos *[]map[string]int) bool {
	// Base case: design reduced to ""
	if len(design) == 0 {
		// Copy the valid combo and store it in allCombos
		validCombo := make(map[string]int)
		for k, v := range combo {
			validCombo[k] = v
		}
		*allCombos = append(*allCombos, validCombo)
		return true
	}

	foundValid := false // Track if at least one valid combination is found

	for i := range availTowels {
		currTowel := availTowels[i]
		c_len := len(currTowel)

		// Skip if the towel is longer than the remaining design
		if c_len > len(design) {
			continue
		}

		// Check if the current towel matches the prefix of the design
		ptrn := regexp.MustCompile(fmt.Sprintf("^%s", currTowel))
		match := ptrn.FindStringIndex(design)
		if match == nil {
			continue
		}

		// Consume the towel and update the combo
		combo[currTowel]++
		remainingDesign := design[c_len:]

		// Recursively solve for the remaining design
		if recur_allcombo(availTowels, remainingDesign, combo, allCombos) {
			foundValid = true
		}

		// Backtrack
		combo[currTowel]--
		if combo[currTowel] == 0 {
			delete(combo, currTowel)
		}
	}

	return foundValid
}

func splitUpDesign(idx []int, design string) []string {
	if len(idx) == 0 {
		return []string{design}
	}

	designs := []string{}
	lastIdx := 0

	s, e := idx[0], idx[1]
	left := design[lastIdx:s]
	designs = append(designs, left)
	lastIdx = e

	right := design[lastIdx:]
	designs = append(designs, right)

	return designs
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
	count := 0
	c := make(chan int)
	for _, d := range designs {
		count++
		go func(avail []string, design string) {
			possible := matchDesign(availTowels, d)
			if possible != nil {
				c <- 1
				fmt.Println(d)
			} else {
				c <- 0
			}
		}(availTowels, d)
	}

	for i := 0; i < count; i++ {
		result += <-c
	}

	return result
}

func Part2(availTowels []string, designs []string) int {
	result := 0
	for _, d := range designs {
		result += findAllCombo(availTowels, d)
	}

	return result
}
