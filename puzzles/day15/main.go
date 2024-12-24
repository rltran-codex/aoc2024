package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Direction int

const (
	Up = iota
	Down
	Right
	Left
)

type LanternfishSchool struct {
	Warehouse        [][]string
	Instructions     []string
	CurrentPosition  Coordinate
	CurrentDirection Direction
	Edges            struct {
		lowerX int
		lowerY int
		upperX int
		upperY int
	}
}

type Coordinate struct {
	X int
	Y int
}

func (ls *LanternfishSchool) gatherGPS() int {
	sum := 0
	for r := 0; r < len(ls.Warehouse); r++ {
		for c := 0; c < len(ls.Warehouse[r]); c++ {
			if ls.Warehouse[r][c] != "O" {
				continue
			}

			sum += (100 * r) + c
		}
	}

	return sum
}

func (ls *LanternfishSchool) performMove() {
	// pop the instruction
	nextMove := utils.PopQueue(&ls.Instructions)
	nextCoor, nextTile := ls.parseMove(nextMove)
	// if the next move contains a #, don't do anything
	switch nextTile {
	case "#":
		return
	case ".":
		ls.Warehouse[nextCoor.Y][nextCoor.X] = "@"
		ls.Warehouse[ls.CurrentPosition.Y][ls.CurrentPosition.X] = "."
		ls.CurrentPosition = nextCoor
	case "O":
		// count how many boxes there are
		// push the boxes
		ls.handleBoxMove()
	}
}

func (ls *LanternfishSchool) handleBoxMove() {
	pushBoxes := func(row int, leftPush bool) {
		leftBound := ls.CurrentPosition.X
		rightBound := ls.CurrentPosition.X
		for {
			if ls.Warehouse[row][rightBound] == "#" || ls.Warehouse[row][rightBound] == "." {
				break
			}
			if leftPush {
				rightBound--
			} else {
				rightBound++
			}
		}
		// make the shift
		if ls.Warehouse[row][rightBound] == "#" {
			return
		}

		r := []string{}
		if leftPush {
			leftBound, rightBound = rightBound, leftBound
			r = append(r, ls.Warehouse[row][leftBound+1:rightBound+1]...)
			r = append(r, ".")
		} else {
			r = append(r, ".")
			r = append(r, ls.Warehouse[row][leftBound:rightBound]...)
		}
		for i, v := range r {
			ls.Warehouse[row][i+leftBound] = v
		}
	}

	// if up or down, rotate the map to make moving boxes easy
	// if left or right move easily find the length of the boxes to move
	switch ls.CurrentDirection {
	case Up:
		ls.Warehouse = utils.Rotate2DSlice(ls.Warehouse, utils.ANTICLOCK)
		ls.findFish()
		pushBoxes(ls.CurrentPosition.Y, true)
		ls.Warehouse = utils.Rotate2DSlice(ls.Warehouse, utils.CLOCKWISE)
		ls.findFish()
	case Down:
		ls.Warehouse = utils.Rotate2DSlice(ls.Warehouse, utils.ANTICLOCK)
		ls.findFish()
		pushBoxes(ls.CurrentPosition.Y, false)
		ls.Warehouse = utils.Rotate2DSlice(ls.Warehouse, utils.CLOCKWISE)
		ls.findFish()
	case Left:
		pushBoxes(ls.CurrentPosition.Y, true)
		ls.findFish()
	case Right:
		pushBoxes(ls.CurrentPosition.Y, false)
		ls.findFish()
	}

}

func (ls *LanternfishSchool) findFish() {
	for r := 0; r < len(ls.Warehouse); r++ {
		for c := 0; c < len(ls.Warehouse[0]); c++ {
			if ls.Warehouse[r][c] == "@" {
				ls.CurrentPosition.X = c
				ls.CurrentPosition.Y = r
				return
			}
		}
	}
}

func (ls *LanternfishSchool) parseMove(dir string) (Coordinate, string) {
	x := 0
	y := 0
	switch dir {
	case "^":
		y--
		ls.CurrentDirection = Up
	case ">":
		x++
		ls.CurrentDirection = Right
	case "v":
		y++
		ls.CurrentDirection = Down
	case "<":
		x--
		ls.CurrentDirection = Left
	default:
		panic("unknown move")
	}

	nextCoordinate := Coordinate{
		X: (ls.CurrentPosition.X + x),
		Y: (ls.CurrentPosition.Y + y),
	}
	nextTile := ls.Warehouse[nextCoordinate.Y][nextCoordinate.X]
	return nextCoordinate, nextTile
}

func main() {
	data := ParsePuzzleInput(false, "day15.txt")
	// main area to display puzzle answers
	fmt.Printf("Part 1: %d\n", Part1(data))
	fmt.Printf("Part 2: %d\n", Part2())
}

func ParsePuzzleInput(sample bool, filename string) *LanternfishSchool {
	puzzle := utils.GetPuzzleInput(filename, sample)
	defer puzzle.Close()

	scn := bufio.NewScanner(puzzle)

	moveInstructions := false

	warehouse := [][]string{}
	instructions := []string{}
	for scn.Scan() {
		line := scn.Text()
		if len(strings.TrimSpace(line)) == 0 {
			moveInstructions = true
		}

		l := strings.Split(line, "")
		if moveInstructions {
			instructions = append(instructions, l...)
		} else {
			warehouse = append(warehouse, l)
		}
	}

	// find the starting position
	currPos := Coordinate{}
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[0]); j++ {
			if warehouse[i][j] == "@" {
				currPos.X = j
				currPos.Y = i
			}
		}
	}

	return &LanternfishSchool{
		Warehouse:       warehouse,
		Instructions:    instructions,
		CurrentPosition: currPos,
		Edges: struct {
			lowerX int
			lowerY int
			upperX int
			upperY int
		}{
			lowerX: 0,
			lowerY: 0,
			upperX: len(warehouse) - 1,
			upperY: len(warehouse[0]) - 1,
		},
	}
}

func Part1(ls *LanternfishSchool) int {
	for len(ls.Instructions) > 0 {
		ls.performMove()
	}

	return ls.gatherGPS()
}
func Part2() int {
	panic("Not yet implemented")
}
