package main

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Coordinate struct {
	X int
	Y int
}

func findShortestPath(grid [][]string, startPoint Coordinate, endPoint Coordinate) ([]Coordinate, int) {
	visited := make([][]bool, len(grid))
	dist := make([][]int, len(grid))
	trace := make([][]*Coordinate, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid))
		dist[i] = make([]int, len(grid))
		trace[i] = make([]*Coordinate, len(grid))
		for j := range dist[i] {
			dist[i][j] = math.MaxInt // set all points to infinity
		}
	}

	// Directions for movement (up, down, left, right)
	// NOTE: by changing the order of the direction, the sample's shortest path now matches the algorithm's shortest path
	directions := []Coordinate{
		{X: 1, Y: 0},  // Right
		{X: 0, Y: 1},  // Down
		{X: -1, Y: 0}, // Left
		{X: 0, Y: -1}, // Up
	}

	// initialize BFS
	queue := []Coordinate{startPoint}
	visited[startPoint.Y][startPoint.X] = true
	dist[startPoint.Y][startPoint.X] = 0

	isValid := func(adj Coordinate) bool {
		return adj.X >= 0 && adj.X < len(grid) &&
			adj.Y >= 0 && adj.Y < len(grid) &&
			!visited[adj.Y][adj.X] && grid[adj.Y][adj.X] != "#"
	}

	buildPath := func(parent [][]*Coordinate, startPoint, endPoint Coordinate) []Coordinate {
		var path []Coordinate
		for current := &endPoint; current != nil; current = parent[current.Y][current.X] {
			path = append([]Coordinate{*current}, path...)
		}
		return path
	}

	for len(queue) > 0 {
		curr := utils.PopQueue(&queue)
		if curr == endPoint {
			return buildPath(trace, startPoint, endPoint), dist[curr.Y][curr.X]
		}

		for _, dir := range directions {
			adj := Coordinate{
				X: curr.X + dir.X,
				Y: curr.Y + dir.Y,
			}

			if !isValid(adj) {
				continue
			}

			visited[adj.Y][adj.X] = true
			dist[adj.Y][adj.X] = dist[curr.Y][curr.X] + 1
			trace[adj.Y][adj.X] = &curr
			queue = append(queue, adj)
		}
	}

	return nil, -1
}

func generateGrid(size int, bytes []Coordinate) [][]string {
	grid := make([][]string, size+1)
	for i := range grid {
		grid[i] = make([]string, size+1)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for _, c := range bytes {
		grid[c.Y][c.X] = "#"
	}

	return grid
}

func main() {
	// main area to display puzzle answers
	input := ParsePuzzleInput(false, "day18.txt")
	fmt.Printf("Part 1: %d\n", Part1(70, input, 1024))
	fmt.Printf("Part 2: %s\n", Part2(70, input, 1024))
}

func ParsePuzzleInput(sample bool, filename string) []Coordinate {
	data := utils.GetPuzzleInput(filename, sample)
	scn := bufio.NewScanner(data)

	bytePos := []Coordinate{}
	for scn.Scan() {
		coor := strings.Split(scn.Text(), ",")
		bytePos = append(bytePos, Coordinate{
			X: utils.Atoi(coor[0]),
			Y: utils.Atoi(coor[1]),
		})
	}

	return bytePos
}

func Part1(gridSize int, bytes []Coordinate, initBytes int) int {
	grid := generateGrid(gridSize, bytes[:initBytes])
	endPt := Coordinate{
		X: gridSize,
		Y: gridSize,
	}
	startPt := Coordinate{
		X: 0,
		Y: 0,
	}

	_, dist := findShortestPath(grid, startPt, endPt)
	return dist
}

func Part2(gridSize int, bytes []Coordinate, initBytes int) string {
	grid := generateGrid(gridSize, bytes[:initBytes])
	endPt := Coordinate{
		X: gridSize,
		Y: gridSize,
	}
	startPt := Coordinate{
		X: 0,
		Y: 0,
	}
	// one at a time, drop bytes.
	path, _ := findShortestPath(grid, startPt, endPt)
	for _, v := range bytes[initBytes:] {
		grid[v.Y][v.X] = "#"
		// if "#" lands on the path, then recalculate the path
		if utils.Index(path, v) != -1 {
			path, _ = findShortestPath(grid, startPt, endPt)
			if path == nil {
				return fmt.Sprintf("%d,%d", v.X, v.Y)
			}
		}
	}

	panic("no bytes that fell block the exit")
}
