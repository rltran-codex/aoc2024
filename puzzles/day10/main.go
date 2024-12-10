package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
	structures "github.com/rltran-codex/aoc-2024-go/utils/structures"
)

func main() {
	graph, startingPts := ParsePuzzleInput()
	// main area to display puzzle answers
	fmt.Print(graph, startingPts)
	fmt.Printf("Part 1: %d\n", Part1())
	fmt.Printf("Part 2: %d\n", Part2())
}

func parseKey(x int, y int) string {
	nx := strconv.Itoa(x)
	ny := strconv.Itoa(y)
	return strings.Join([]string{nx, ny}, ";")
}

func ParsePuzzleInput() (structures.Graph, map[string]*structures.GraphNode) {
	// function to parse the puzzle input from file
	lavamap := utils.Get2DPuzzleInput("day10.txt", true)
	maxRow := len(lavamap)
	maxCol := len(lavamap[0])

	ngraph := structures.Graph{
		Nodes: make(map[string]*structures.GraphNode),
		Size:  0,
	}

	startingPts := make(map[string]*structures.GraphNode)

	for i := range lavamap {
		for j := range lavamap[i] {
			key := parseKey(i, j)

			currNode, err := ngraph.GetGNode(key)
			if err != nil {
				v, _ := strconv.Atoi(lavamap[i][j])
				currNode = &structures.GraphNode{
					Id:    key,
					Value: v,
				}
				if v == 0 {
					startingPts[key] = currNode
				}
			}

			pos := []struct {
				Row int
				Col int
			}{
				{i - 1, j}, // up
				{i + 1, j}, // down
				{i, j - 1}, // left
				{i, j + 1}, // right
			}
			var adjNodes []*structures.GraphNode
			for _, p := range pos {
				if p.Row < 0 || p.Row >= maxRow || p.Col < 0 || p.Col >= maxCol {
					continue
				}
				key := parseKey(p.Row, p.Col)
				adjNode, err := ngraph.GetGNode(key)
				if err != nil {
					v, _ := strconv.Atoi(lavamap[p.Row][p.Col])
					adjNode = &structures.GraphNode{
						Id:    key,
						Value: v,
					}
				}
				// queue node to be added
				adjNodes = append(adjNodes, adjNode)
			}

			// link graph node
			ngraph.AddGNode(currNode, adjNodes...)
		}
	}

	return ngraph, startingPts
}

func Part1() int {
	// for each starting point, find the paths that reach to 9.
	// Note: a starting point can have more than 1 path each

}
func Part2() int {
	panic("Not yet implemented")
}
