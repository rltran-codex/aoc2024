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
	Rotation         []Direction         // the rotation cycle (queue)
	Visited          map[string]PathNode // map of the distinct tiles the guard visited linked to  recorded direction they were facing
	PatrolArea       *[][]string
	Path             *PathNode
}

type PathNode struct {
	X         int
	Y         int
	FacingDir Direction
	Next      *PathNode
}

func (g *Guard) StepsTaken() int {
	return len(g.Visited)
}

func (g *Guard) PatrolAreaWithRotation(steps *int) {
	// base case 1: if next step is out of bounds, guard is done patrolling
	x, y := g.CurrentPosition[0], g.CurrentPosition[1]
	// increment distinct steps
	if (*g.PatrolArea)[x][y] == "." || (*g.PatrolArea)[x][y] == "^" {
		*steps++
	}

	// mark it as stepped through
	(*g.PatrolArea)[x][y] = "X"
	if x-1 < 0 {
		return
	}

	// if the next step is a #, rotate map counter clockwise to simulate moving right
	if (*g.PatrolArea)[x-1][y] == "#" {
		*g.PatrolArea = utils.Rotate2DSlice(*g.PatrolArea, utils.ANTICLOCK)

		// update curr position since map was transposed
		n := len(*g.PatrolArea)
		nx, ny := n-1-y, x
		x, y = nx, ny
	}

	g.CurrentPosition[0] = x - 1
	g.CurrentPosition[1] = y
	g.PatrolAreaWithRotation(steps)
}

func (g *Guard) GetPatrolPath() {
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

		// add this tile to visited
		idx := parseIndices(x, y)
		_, ok := g.Visited[idx]
		if !ok {
			g.Visited[idx] = *currNode
		}

		// guard's next step is out of bounds, loop done
		if nx < 0 || nx >= maxHeight || ny < 0 || ny >= maxWidth {
			return
		}

		if (*g.PatrolArea)[nx][ny] == "#" {
			g.CurrentDirection = utils.PopAndRequeue(&g.Rotation)
			nx, ny = calcNextStep(x, y, g.CurrentDirection)
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

func parseIndices(x int, y int) string {
	xstr := strconv.Itoa(x)
	ystr := strconv.Itoa(y)
	return strings.Join([]string{xstr, ystr}, ";")
}

// note: need to handle when reaching the head and the next step is not an already visited tile
func (g *Guard) DetectCycle() bool {
	panic("TODO")
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
	data1, start1 := ParsePuzzleInput()
	data2, start2 := ParsePuzzleInput()

	fmt.Printf("Part 1: %d\n", Part1(data1, start1))
	fmt.Printf("Part 2: %d\n", Part2(data2, start2))
}

func Part1(data [][]string, start []int) int {
	results := 0
	guard := Guard{
		PatrolArea:      &data,
		CurrentPosition: start,
	}
	// patrolAreaWithRotation(&data, start, &results)
	guard.PatrolAreaWithRotation(&results)
	return results
}

func Part2(data [][]string, start []int) int {
	guard := Guard{
		CurrentPosition:  []int{start[0], start[1]},
		CurrentDirection: DefaultDir,
		Rotation:         DefaultRot,
		Visited:          make(map[string]PathNode),
		PatrolArea:       &data,
	}

	guard.GetPatrolPath()

	return 0
}

func ParsePuzzleInput() ([][]string, []int) {
	file := utils.GetPuzzleInput("day6.txt", true)
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
