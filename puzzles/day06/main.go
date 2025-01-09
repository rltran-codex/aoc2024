package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Direction string

const (
	UP    = "UP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

var DefaultRot = []Direction{RIGHT, DOWN, LEFT, UP}
var DefaultDir Direction = UP

type Guard struct {
	CurrentPosition  []int
	CurrentDirection Direction
	Rotation         []Direction     // the rotation cycle (queue)
	Visited          map[string]bool // map of the distinct tiles the guard visited linked to recorded direction they were facing
	PatrolArea       *[][]string
	Path             *PathNode
}

type PathNode struct {
	X         int
	Y         int
	FacingDir Direction
	Next      *PathNode
}

func (g *Guard) DistinctLocations() int {
	results := make(map[string]struct{})
	for curr := g.Path; curr != nil; {
		key := utils.ParseKey(curr.X, curr.Y)
		results[key] = struct{}{}
		curr = curr.Next
	}

	return len(results)
}

func (g *Guard) GetPatrolPath(detectCycle bool) bool {
	maxHeight := len(*g.PatrolArea)
	maxWidth := len((*g.PatrolArea)[0])

	// current position is the starting Path
	g.Path = &PathNode{
		X:         g.CurrentPosition[0],
		Y:         g.CurrentPosition[1],
		FacingDir: g.CurrentDirection,
	}

	currNode := g.Path
	for {
		x, y := currNode.X, currNode.Y
		nx, ny := calcNextStep(x, y, g.CurrentDirection)

		// guard's next step is out of bounds, loop done
		if nx < 0 || nx >= maxHeight || ny < 0 || ny >= maxWidth {
			return false
		}

		if (*g.PatrolArea)[nx][ny] == "#" {
			// turn guard
			g.CurrentDirection = utils.PopAndRequeue(&g.Rotation)
			currNode.FacingDir = g.CurrentDirection
			continue
		}

		if detectCycle {
			key := parseIndices(x, y, g.CurrentDirection)
			if g.Visited[key] {
				return true
			}
			g.Visited[key] = true
		}

		nextNode := PathNode{
			X:         nx,
			Y:         ny,
			FacingDir: g.CurrentDirection,
		}

		currNode.Next = &nextNode
		currNode = &nextNode
	}
}

func parseIndices(x int, y int, dir Direction) string {
	xstr := strconv.Itoa(x)
	ystr := strconv.Itoa(y)
	return strings.Join([]string{xstr, ystr, string(dir)}, ";")
}

func calcNextStep(x int, y int, currDir Direction) (int, int) {
	var nextX, nextY int
	switch currDir {
	case UP:
		nextX = x - 1
		nextY = y
	case DOWN:
		nextX = x + 1
		nextY = y
	case LEFT:
		nextX = x
		nextY = y - 1
	case RIGHT:
		nextX = x
		nextY = y + 1
	}

	return nextX, nextY
}

func main() {
	// main area to display puzzle answers
	data1, start1 := ParsePuzzleInput(false)
	data2, start2 := ParsePuzzleInput(false)

	fmt.Printf("Part 1: %d\n", Part1(data1, start1))
	fmt.Printf("Part 2: %d\n", Part2(data2, start2))
}

func Part1(data [][]string, start []int) int {
	guard := Guard{
		CurrentPosition:  []int{start[0], start[1]},
		CurrentDirection: DefaultDir,
		Rotation:         DefaultRot,
		PatrolArea:       &data,
	}

	// conduct part 1, get the patrol path
	guard.GetPatrolPath(false)
	return guard.DistinctLocations()
}

func Part2(data [][]string, start []int) int {
	guard := Guard{
		CurrentPosition:  []int{start[0], start[1]},
		CurrentDirection: DefaultDir,
		Rotation:         DefaultRot,
		PatrolArea:       &data,
	}

	// conduct part 1, get the patrol path
	guard.GetPatrolPath(false)
	maxHeight := len(data)
	maxWidth := len(data[0])
	validObstacles := make(map[string]bool)
	for curr := guard.Path; curr != nil; curr = curr.Next {
		// create a new guard and place obstacles in front of current node
		ng := Guard{
			CurrentPosition:  []int{start[0], start[1]},
			CurrentDirection: DefaultDir,
			Rotation:         DefaultRot,
			PatrolArea:       deepCopy2DArray(data),
			Visited:          make(map[string]bool),
		}

		nx, ny := calcNextStep(curr.X, curr.Y, curr.FacingDir)
		key := utils.ParseKey(nx, ny)
		if validObstacles[key] {
			continue
		}

		// if next step is in bounds, place obstacle in front of it
		if nx < 0 || nx >= maxHeight || ny < 0 || ny >= maxWidth {
			continue
		}
		(*ng.PatrolArea)[nx][ny] = "#"
		if ng.GetPatrolPath(true) {
			validObstacles[key] = true
		}
	}
	return len(validObstacles)
}

func ParsePuzzleInput(sample bool) ([][]string, []int) {
	file := utils.GetPuzzleInput("day6.txt", sample)
	defer file.Close()

	scn := bufio.NewScanner(file)
	var data [][]string
	startingPosition := make([]int, 2)

	row := 0
	for scn.Scan() {
		line := scn.Text()
		i := strings.Index(line, "^")
		if i != -1 {
			startingPosition[0] = row
			startingPosition[1] = i
		}
		data = append(data, strings.Split(line, ""))

		row++
	}

	return data, startingPosition
}

func deepCopy2DArray(original [][]string) *[][]string {
	n := len(original)
	m := len(original[0])
	duplicate := make([][]string, n)
	data := make([]string, n*m)
	for i := range original {
		start := i * m
		end := start + m
		duplicate[i] = data[start:end:end]
		copy(duplicate[i], original[i])
	}

	return &duplicate
}
