package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rltran-codex/aoc-2024-go/utils"
	structures "github.com/rltran-codex/aoc-2024-go/utils/structures"
)

func main() {
	_, startingPts := ParsePuzzleInput()
	// main area to display puzzle answers
	fmt.Printf("Part 1: %d\n", Part1(startingPts))
	fmt.Printf("Part 2: %d\n", Part2(startingPts))
}

func Part1(startingPts map[string]*structures.GraphNode) int {
	results := 0
	// for each starting point, find the paths that reach to 9.
	// Note: a starting point can have more than 1 path each
	for _, v := range startingPts {
		visited9s := make(map[string]*structures.GraphNode)
		r := findHikingPath(v, visited9s, false)
		results += r
	}

	return results
}

func Part2(startingPts map[string]*structures.GraphNode) int {
	results := 0

	for _, v := range startingPts {
		visited9s := make(map[string]*structures.GraphNode)
		r := findHikingPath(v, visited9s, true)
		results += r
	}

	return results
}

func findHikingPath(currNode *structures.GraphNode, visited9s map[string]*structures.GraphNode, countPaths bool) int {
	// base case 1: current node value is 9
	curr, _ := currNode.Value.(int)
	if curr == 9 {
		if countPaths {
			return 1
		}

		_, ok := visited9s[currNode.Id]
		if !ok {
			visited9s[currNode.Id] = currNode
			return 1
		} else {
			return 0
		}
	}

	// traverse nodes that only increment by 1
	n := 0
	for _, v := range currNode.Adj {
		adj, _ := v.Value.(int)
		diff := adj - curr
		if diff == 1 {
			n += findHikingPath(v, visited9s, countPaths)
		}
	}
	return n
}

func parseKey(x int, y int) string {
	nx := strconv.Itoa(x)
	ny := strconv.Itoa(y)
	return strings.Join([]string{nx, ny}, ";")
}

func ParsePuzzleInput() (structures.Graph, map[string]*structures.GraphNode) {
	// function to parse the puzzle input from file
	lavamap := utils.Get2DPuzzleInput("day10.txt", false)
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
			}
			if lavamap[i][j] == "0" {
				startingPts[key] = currNode
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
