package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func main() {
	// main area to display puzzle answers
	pageOrdr, printPage := ParsePuzzleInput()
	result, fSection := Part1(pageOrdr, printPage)
	fmt.Printf("Part 1: %d\n", result)
	result = Part2(pageOrdr, fSection)
	fmt.Printf("Part 2: %d\n", result)
}

func Part2(pageOrdr map[string][]string, failedSections []string) int {
	panic("not yet done")
}

func Part1(pageOrdr map[string][]string, prodPage []string) (int, []string) {
	results := 0
	r := regexp.MustCompile(`\d+`)
	var failures []string

	for _, v := range prodPage {
		nums := r.FindAllString(v, -1)

		correct := true
		offset := 0
		for i := 0; i < len(nums)-1; i++ {
			rules, ok := pageOrdr[nums[i]]
			if !ok {
				continue
			}

			pattern := strings.Join(rules, "|")
			orderReg := regexp.MustCompile(pattern)

			// build the rest of the string
			offset += len(nums[i]) + 1
			cmd := []byte(v[offset:])
			match := orderReg.Find(cmd)
			if len(match) > 0 {
				// this section failed, page ordering rule violated.
				fmt.Println(v, "is incorrect")
				correct = false
				failures = append(failures, v)
				break
			}
		}

		if correct {
			middlePage := nums[len(nums)/2]
			n, _ := strconv.Atoi(middlePage)
			results += n
		}
	}
	return results, failures
}

// Parse the puzzle input.
//
//	Page Ordering Rule is stored in a map and in this manner: after -> [before...]
//	by storing it like this, the puzzle can be solved using regex to check if
//	any numbers that are supposed to be printed before are found ahead
func ParsePuzzleInput() (map[string][]string, []string) {
	pageOrdr := make(map[string][]string)
	var pageProd []string

	file := utils.GetPuzzleInput("day5.txt", false)
	defer file.Close()

	scn := bufio.NewScanner(file)
	for scn.Scan() {
		line := scn.Text()
		if strings.Contains(line, "|") {
			nums := strings.Split(line, "|")
			before := nums[0]
			after := nums[1]
			pageOrdr[after] = append(pageOrdr[after], before)
		} else if strings.Contains(line, ",") {
			pageProd = append(pageProd, strings.TrimSpace(line))
		}
	}

	return pageOrdr, pageProd
}
