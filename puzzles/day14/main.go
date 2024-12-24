package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

// need to find where the robots are after 100 seconds, then split grid into quadrants.
var MinHeight = 0
var MinWidth = 0
var MaxWidth = 101
var MaxHeight = 103

type Robot struct {
	Id       int
	Velocity struct {
		X int
		Y int
	}
	Position struct {
		X int
		Y int
	}
}

func (rbt *Robot) moveBot() {
	rbt.Position.X += rbt.Velocity.X
	rbt.Position.Y += rbt.Velocity.Y

	// check if need to teleport robot on the X and Y axis
	if rbt.Position.X < 0 {
		rbt.Position.X = rbt.Position.X + MaxWidth
	} else if rbt.Position.X >= MaxWidth {
		rbt.Position.X = rbt.Position.X % MaxWidth
	}

	if rbt.Position.Y < 0 {
		rbt.Position.Y = rbt.Position.Y + MaxHeight
	} else if rbt.Position.Y >= MaxHeight {
		rbt.Position.Y = rbt.Position.Y % MaxHeight
	}
}

func (rbt *Robot) moveBotN(n int) {
	for i := 0; i < n; i++ {
		rbt.moveBot()
	}
}

func splitGridAndCount(rbts []*Robot) map[string]int {
	// split the grid and iterate
	// through all robots to see which quadrant they belong to
	midBoundX := MaxWidth / 2
	midBoundY := MaxHeight / 2
	// NOTE: Robots that are exactly in the middle (horizontally or vertically) don't count as being in any quadrant
	quadrants := map[string]struct {
		lowerX int
		upperX int
		lowerY int
		upperY int
	}{
		"top_left":  {lowerX: 0, upperX: midBoundX - 1, lowerY: 0, upperY: midBoundY - 1},
		"top_right": {lowerX: midBoundX + 1, upperX: MaxWidth - 1, lowerY: 0, upperY: midBoundY - 1},
		"btm_left":  {lowerX: 0, upperX: midBoundX - 1, lowerY: midBoundY + 1, upperY: MaxHeight - 1},
		"btm_right": {lowerX: midBoundX + 1, upperX: MaxWidth - 1, lowerY: midBoundY + 1, upperY: MaxHeight - 1},
	}

	rbtQuad := map[string]int{
		"top_left":  0,
		"top_right": 0,
		"btm_left":  0,
		"btm_right": 0,
	}

	for _, rbt := range rbts {
		for quad, bounds := range quadrants {
			if rbt.Position.X >= bounds.lowerX && rbt.Position.X <= bounds.upperX &&
				rbt.Position.Y >= bounds.lowerY && rbt.Position.Y <= bounds.upperY {
				rbtQuad[quad]++
				break
			}
		}
	}

	return rbtQuad
}

// Q1: x < MidWidth and y < MidHeight
// Q2: x >= MidWidth and y < MidHeight
// Q3: x < MidWidth and y >= MidHeight
// Q4: x >= MidWidth and y >= MidHeight

func main() {
	// main area to display puzzle answers
	data := ParsePuzzleInput(false)
	fmt.Printf("Part 1: %d\n", Part1(data))
	data = ParsePuzzleInput(false)
	fmt.Printf("Part 2: %d\n", Part2(data))
}

func ParsePuzzleInput(sample bool) []*Robot {
	data := utils.GetPuzzleInput("day14.txt", sample)
	defer data.Close()
	r := regexp.MustCompile(`[p|v]=(-?[0-9]+),(-?[0-9]+)`)

	scn := bufio.NewScanner(data)
	robots := []*Robot{}
	id := 1
	for scn.Scan() {
		line := scn.Text()
		match := r.FindAllStringSubmatch(line, 2)
		robots = append(robots, &Robot{
			Id: id,
			Velocity: struct {
				X int
				Y int
			}{
				X: utils.Atoi(match[1][1]),
				Y: utils.Atoi(match[1][2]),
			},
			Position: struct {
				X int
				Y int
			}{
				X: utils.Atoi(match[0][1]),
				Y: utils.Atoi(match[0][2]),
			},
		})

		id++
	}

	return robots
}

func Part1(rbts []*Robot) int {
	var wg sync.WaitGroup
	for _, r := range rbts {
		wg.Add(1)
		func() {
			defer wg.Done()
			r.moveBotN(100)
		}()
	}
	wg.Wait()

	result := 1
	quad := splitGridAndCount(rbts)
	for _, v := range quad {
		result *= v
	}

	return result
}

func Part2(rbts []*Robot) int {
	rbtPositions := make(map[string]*Robot)
	sec := 0
	for len(rbtPositions) != len(rbts) {
		rbtPositions = make(map[string]*Robot)
		for _, v := range rbts {
			v.moveBot()
			k := parseKey(v.Position.X, v.Position.Y)
			rbtPositions[k] = v
		}
		sec++
	}

	drawGrid(rbts)
	return sec
}

func parseKey(x int, y int) string {
	nx := strconv.Itoa(x)
	ny := strconv.Itoa(y)
	return strings.Join([]string{nx, ny}, ";")
}

// for debugging
func drawGrid(rbts []*Robot) [][]int {
	grid := make([][]int, MaxHeight)
	for i := range grid {
		grid[i] = make([]int, MaxWidth)
	}

	for _, r := range rbts {
		grid[r.Position.Y][r.Position.X]++
	}

	for i := range grid {
		fmt.Println(grid[i])
	}
	return grid
}
