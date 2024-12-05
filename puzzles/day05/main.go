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
	results := 0
	c := make(chan int)

	for _, section := range failedSections {
		go swapUntilValid(pageOrdr, section, c)
	}

	for range failedSections {
		results += <-c
	}
	close(c)
	return results
}

func swapUntilValid(pageOrdr map[string][]string, section string, c chan int) {
	r := regexp.MustCompile(`\d+`)
	pages := r.FindAllString(section, -1)
	offset := 0

	for i := 0; i < len(pages)-1; i++ {
		currPage := pages[i]
		offset += len(currPage) + 1
		str := section[offset:]
		update := []byte(str)
		ok, fail := validate(currPage, update, pageOrdr)
		if !ok {
			// swap and reiterate
			n := string(update[fail[0]:fail[1]])
			relativeIdx := utils.Index(pages[i+1:], n)
			targetIdx := i + relativeIdx + 1
			pages[i], pages[targetIdx] = pages[targetIdx], pages[i]

			// restart loop... this isnt optimal but it gets it done...
			offset = 0
			i = -1
			section = strings.Join(pages, ",")
		}
	}

	middlePage := pages[len(pages)/2]
	n, _ := strconv.Atoi(middlePage)
	c <- n
}

func Part1(pageOrdr map[string][]string, prodPage []string) (int, []string) {
	results := 0
	r := regexp.MustCompile(`\d+`)
	var failures []string

	for _, v := range prodPage {
		pages := r.FindAllString(v, -1)

		correct := true
		offset := 0
		for i := 0; i < len(pages)-1; i++ {
			offset += len(pages[i]) + 1
			update := []byte(v[offset:])

			if ok, _ := validate(pages[i], update, pageOrdr); !ok {
				failures = append(failures, v)
				correct = false
				break
			}
		}

		if correct {
			middlePage := pages[len(pages)/2]
			n, _ := strconv.Atoi(middlePage)
			results += n
		}
	}
	return results, failures
}

func validate(currPage string, update []byte, pageOrdr map[string][]string) (bool, []int) {
	rules, ok := pageOrdr[currPage]
	if !ok {
		return true, []int{}
	}

	pattern := strings.Join(rules, "|")
	orderReg := regexp.MustCompile(pattern)

	// build the rest of the string
	match := orderReg.FindIndex(update)
	return len(match) == 0, match
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
			pages := strings.Split(line, "|")
			before := pages[0]
			after := pages[1]
			pageOrdr[after] = append(pageOrdr[after], before)
		} else if strings.Contains(line, ",") {
			pageProd = append(pageProd, strings.TrimSpace(line))
		}
	}

	return pageOrdr, pageProd
}
