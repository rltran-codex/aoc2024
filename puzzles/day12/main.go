package main

import (
	"fmt"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

// {x, y}
var directions = [][2]int{
	{0, -1}, // up
	{1, 0},  // right
	{0, 1},  // down
	{-1, 0}, // left
}

func main() {
	garden := ParsePuzzleInput(false, "day12.txt")
	// main area to display puzzle answers
	fmt.Printf("Part 1: %d\n", Part1(garden))
	fmt.Printf("Part 2: %d\n", Part2(garden))
}

func ParsePuzzleInput(sample bool, filename string) [][]string {
	// function to parse the puzzle input from file
	return utils.Get2DPuzzleInput(filename, sample)
}

func Part1(garden [][]string) (cost int) {
	n, m := len(garden), len(garden[0])
	visited := make([][]bool, n)
	for i := range garden {
		visited[i] = make([]bool, m)
	}

	for y := range garden {
		for x := range garden {
			if visited[y][x] {
				continue
			}

			id, currPos := garden[y][x], [2]int{x, y}
			area, perimeter, _ := fence(garden, visited, id, currPos)
			cost += (area * perimeter)
		}
	}
	return cost
}

func Part2(garden [][]string) (cost int) {
	n, m := len(garden), len(garden[0])
	visited := make([][]bool, n)
	for i := range garden {
		visited[i] = make([]bool, m)
	}

	for y := range garden {
		for x := range garden {
			if visited[y][x] {
				continue
			}

			id, currPos := garden[y][x], [2]int{x, y}
			area, _, sides := fence(garden, visited, id, currPos)
			// fmt.Printf("%s: %d %d\n", garden[y][x], area, sides)
			cost += (area * sides)
		}
	}
	return cost
}

// dfs traversal
func fence(garden [][]string, visited [][]bool, id string, position [2]int) (area int, perimeter int, corners int) {
	x, y := position[0], position[1]
	// Base Case 1: out of bounds or region != id, increment perimeter only
	if x < 0 || x >= len(garden) || y < 0 || y >= len(garden[0]) || garden[y][x] != id {
		return 0, 1, 0
	}
	// Base Case 2: if visited, do nothing
	if visited[y][x] {
		return 0, 0, 0
	}

	visited[y][x] = true
	area += 1
	// check corners for current tile
	corners += cornerCount(garden, position)
	for _, v := range directions[:4] {
		nx, ny := x+v[0], y+v[1]
		a, p, c := fence(garden, visited, id, [2]int{nx, ny})
		area += a
		perimeter += p
		corners += c
	}

	return area, perimeter, corners
}

/*
idea here is to rotate the L shaped inspection area around the tile point
and check two cases for corners.
for example:
1st
|*|*| |
|*|x| |
| | | |

2nd
| |*|*|
| |x|*|
| | | |

3rd
| | | |
| |x|*|
| |*|*|

4th
| | | |
|*|x| |
|*|*| |

where * is the tile to inspect and x is the current tile (pos)
In each rotation, check for these case (use 1st case as reference for this):
 1. if upper and left is are not in bounds or same flower, then corner found
 2. if upper and left are true, but left is false, then corner found
*/
func cornerCount(garden [][]string, pos [2]int) (corners int) {
	x, y := pos[0], pos[1]
	flower := garden[y][x]
	maxWidth, maxHeight := len(garden[0]), len(garden)

	dir := []*[2]int{
		{-1, -1}, // up left
		{-1, 0},  // up
		{0, -1},  // left
	}

	rotate := func() {
		for _, d := range dir {
			t := d[0]
			d[0] = d[1]
			d[1] = -1 * t
		}
	}
	inBounds := func(d [2]int) bool {
		i, j := d[0], d[1]
		return j >= 0 && j < maxHeight && i >= 0 && i < maxWidth
	}

	for i := 0; i < 4; i++ {
		UL := [2]int{x + dir[0][0], y + dir[0][1]}
		UR := [2]int{x + dir[1][0], y + dir[1][1]}
		LL := [2]int{x + dir[2][0], y + dir[2][1]}
		ul := inBounds(UL) && garden[UL[1]][UL[0]] == flower
		ur := inBounds(UR) && garden[UR[1]][UR[0]] == flower
		ll := inBounds(LL) && garden[LL[1]][LL[0]] == flower

		if !ur && !ll {
			corners++
		}
		if ur && ll && !ul {
			corners++
		}

		rotate()
	}

	return corners
}
