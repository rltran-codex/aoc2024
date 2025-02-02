package main

import (
	"fmt"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

var memo map[string]int = make(map[string]int)

var directions = [][]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func main() {
	// main area to display puzzle answers
	t, s := ParsePuzzleInput(false, "day20.txt")
	fmt.Printf("Part 1: %d\n", Part1(t, s[:]))
	fmt.Printf("Part 2: %d\n", Part2())
}

// Note: for part 2,use a counter as max number of # we can phase through (?)
func dfs(before []int, start []int, startDist int, racetrack [][]string) (n int, distMap [][]int) {
	key := fmt.Sprintf("%d:%d-%d", start[0], start[1], startDist)
	if n, ok := memo[key]; ok {
		return n, nil
	}
	defer func() {
		memo[key] = n
	}()

	height := len(racetrack)
	width := len(racetrack[0])
	visited := make(map[[2]int]bool)
	dist := make([][]int, height)
	for i := range racetrack {
		dist[i] = make([]int, width)
	}

	dist[start[0]][start[1]] = startDist
	var dfsRecur func(node [2]int) int
	dfsRecur = func(node [2]int) int {
		x := node[1]
		y := node[0]
		cVal := racetrack[y][x]

		// base case: cycle found
		if before != nil && node == [2]int(before) {
			return -1
		}

		// base case: reached end
		if cVal == "E" {
			return dist[y][x]
		}

		visited[node] = true
		bestDist := -1

		for _, adj := range directions {
			nx, ny := (x + adj[1]), (y + adj[0])
			key := [2]int{ny, nx}
			if visited[key] ||
				(inBounds(nx, ny, width, height) && racetrack[ny][nx] == "#") {
				continue
			}

			dist[ny][nx] = dist[y][x] + 1
			foundDist := dfsRecur([2]int{ny, nx})
			// if E was found
			if foundDist != -1 {
				if bestDist == -1 || foundDist < bestDist {
					bestDist = foundDist
				}
			}
		}

		// failed to find "E", back track
		visited[node] = false
		return bestDist
	}

	return dfsRecur([2]int(start)), dist
}

// Note: for part 2, when calling dfs(curr, cheatStart, (initMap[y][x] + 2), racetrack), do not
// care about whether or not the next step of a wall is "." and
func traverse(start []int, racetrack [][]string) map[int]int {
	width := len(racetrack[0])
	height := len(racetrack)

	visited := make(map[[2]int]bool)
	cheats := make(map[int]int)
	queue := [][]int{start}

	initDist, initMap := dfs(nil, start, 0, racetrack)

	// loop while queue is not empty
	for len(queue) > 0 {
		curr := utils.PopQueue(&queue)
		x := curr[1]
		y := curr[0]

		visited[[2]int(curr)] = true

		for _, adj := range directions {
			nx, ny := (x + adj[1]), (y + adj[0])
			key := [2]int{ny, nx}
			val := racetrack[ny][nx]

			if !inBounds(nx, ny, width, height) {
				continue
			}

			if val == "#" {
				// use dfs to find path to "E"
				if checkWallCheat(nx+adj[1], ny+adj[0], racetrack) {
					cheatStart := []int{ny + adj[0], nx + adj[1]}
					cheatDist, _ := dfs(curr, cheatStart, (initMap[y][x] + 2), racetrack)
					if cheatDist != -1 {
						cheats[(initDist-cheatDist)]++
					}
				}
			} else {
				if visited[key] {
					continue
				}

				queue = append(queue, key[:])
			}
		}
	}

	return cheats
}

func inBounds(x, y, width, height int) bool {
	return x >= 0 && x < width && y >= 0 && y < height
}

func checkWallCheat(x, y int, racetrack [][]string) bool {
	if y < 0 || y >= len(racetrack) || x < 0 || x >= len(racetrack[y]) {
		return false
	}

	val := racetrack[y][x]
	return val == "E" || val == "."
}

// Open "day20.txt" and get
func ParsePuzzleInput(sample bool, filename string) ([][]string, [2]int) {
	track := utils.Get2DPuzzleInput(filename, sample)
	start := make([]int, 2)
	end := make([]int, 2)
	for y := range track {
		for x, v := range track[y] {
			switch v {
			case "S":
				start[0] = y
				start[1] = x
			case "E":
				end[0] = y
				end[1] = x
			}
		}
	}

	return track, [2]int(start)
}

func Part1(track [][]string, start []int) int {
	cheats := traverse(start, track)
	result := 0
	for i := range cheats {
		if i >= 100 {
			result += cheats[i]
		}
	}
	return result
}
func Part2() int {
	panic("Not yet implemented")
}
