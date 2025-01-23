package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type Direction int

const (
	North = iota
	South
	East
	West
)

type Maze struct {
	Start    *MazeTile
	End      *MazeTile
	Tiles    map[string]*MazeTile
	BestPath *[]*MazeTile
}

type MazeTile struct {
	Val   string
	Pos   []int
	Prev  *MazeTile
	Prevs []*MazeTile

	Dist int
	N    *MazeTile
	S    *MazeTile
	E    *MazeTile
	W    *MazeTile
}

type TileState struct {
	Tile *MazeTile
	Dir  Direction
	Cost int
}

func (mt *MazeTile) linkTiles(adj *MazeTile, d Direction) {
	switch d {
	case North:
		mt.N = adj
		adj.S = mt
	case South:
		mt.S = adj
		adj.N = mt
	case East:
		mt.E = adj
		adj.W = mt
	case West:
		mt.W = adj
		adj.E = mt
	}
}

func (mt *MazeTile) getAdj(currDir Direction) []TileState {
	switch currDir {
	case North:
		return []TileState{
			{Tile: mt.N, Dir: North, Cost: 0},
			{Tile: mt.E, Dir: East, Cost: 1000},
			{Tile: mt.W, Dir: West, Cost: 1000},
			{Tile: mt.S, Dir: West, Cost: 2000},
		}
	case South:
		return []TileState{
			{Tile: mt.S, Dir: South, Cost: 0},
			{Tile: mt.E, Dir: East, Cost: 1000},
			{Tile: mt.W, Dir: West, Cost: 1000},
			{Tile: mt.N, Dir: West, Cost: 2000},
		}
	case East:
		return []TileState{
			{Tile: mt.E, Dir: East, Cost: 0},
			{Tile: mt.N, Dir: North, Cost: 1000},
			{Tile: mt.S, Dir: South, Cost: 1000},
			{Tile: mt.W, Dir: South, Cost: 2000},
		}
	case West:
		return []TileState{
			{Tile: mt.W, Dir: West, Cost: 0},
			{Tile: mt.N, Dir: North, Cost: 1000},
			{Tile: mt.S, Dir: South, Cost: 1000},
			{Tile: mt.E, Dir: South, Cost: 2000},
		}
	default:
		panic("unknown direction")
	}
}

func (m *Maze) findShortestPath() int {
	visited := make(map[*MazeTile]bool)
	queue := []*TileState{}
	// set all distances of vertices to infinity
	for _, n := range m.Tiles {
		n.Dist = math.MaxInt
	}
	m.Start.Dist = 0
	queue = append(queue, &TileState{
		Tile: m.Start,
		Dir:  East,
	})

	for len(queue) > 0 {
		sort.Slice(queue, func(i, j int) bool {
			a := queue[i]
			b := queue[j]
			return a.Tile.Dist < b.Tile.Dist
		})

		// pop the vertex with the shortest distance in the queue
		curr := utils.PopQueue(&queue)
		currNode := curr.Tile
		currDir := curr.Dir
		if visited[currNode] {
			continue
		}
		for _, adj := range currNode.getAdj(currDir) {
			adjNode := adj.Tile
			if adjNode.Val == "#" || visited[adjNode] {
				continue
			}

			nDist := currNode.Dist + 1 + adj.Cost
			if adjNode.Dist > nDist {
				adjNode.Dist = nDist
				adjNode.Prev = currNode
			}
			// add to queue
			queue = append(queue, &adj)
		}

		visited[currNode] = true
	}

	return m.End.Dist
}

func (m *Maze) findBestSpots() {
	panic("TODO")
	// build best path
	// find shortest path without turn cost and reference prevs
}

func (m *Maze) getBestPath() []*MazeTile {
	if len(*m.BestPath) > 0 {
		return *m.BestPath
	}

	bestPath := []*MazeTile{}
	for node := m.End; node != nil; node = node.Prev {
		bestPath = append([]*MazeTile{node}, bestPath...)
	}
	m.BestPath = &bestPath
	return *m.BestPath
}

func main() {
	// main area to display puzzle answers
	m := ParsePuzzleInput(false, "day16.txt")
	fmt.Printf("Part 1: %d\n", Part1(m))
	fmt.Printf("Part 2: %d\n", Part2(m))
}

func ParsePuzzleInput(sample bool, filename string) *Maze {
	maze := utils.Get2DPuzzleInput(filename, sample)

	tiles := make(map[string]*MazeTile)
	var sPoint *MazeTile
	var ePoint *MazeTile
	maxR := len(maze)
	maxC := len(maze[0])

	for r := 0; r < maxR; r++ {
		for c := 0; c < maxC; c++ {
			k := utils.ParseKey(r, c)

			ct := tiles[k]
			if ct == nil { // tile not yet touched, add to map
				ct = &MazeTile{Val: maze[r][c], Pos: []int{r, c}}
				tiles[k] = ct
			}

			if ct.Val == "S" {
				sPoint = ct
			} else if ct.Val == "E" {
				ePoint = ct
			}
			adjTiles := []struct {
				R int
				C int
				D Direction
			}{
				{R: r - 1, C: c, D: North},
				{R: r + 1, C: c, D: South},
				{R: r, C: c - 1, D: West},
				{R: r, C: c + 1, D: East},
			}
			for _, v := range adjTiles {
				if v.R < 0 || v.R >= maxR || v.C < 0 || v.C >= maxC {
					continue
				}

				k := utils.ParseKey(v.R, v.C)
				nt := tiles[k]
				if nt == nil {
					nt = &MazeTile{Val: maze[v.R][v.C], Pos: []int{v.R, v.C}}
					tiles[k] = nt
				}
				ct.linkTiles(nt, v.D)
			}
		}
	}

	return &Maze{
		Start: sPoint,
		End:   ePoint,
		Tiles: tiles,
	}
}

func Part1(m *Maze) int {
	return m.findShortestPath()
}
func Part2(m *Maze) int {
	panic("Not yet implemented")
}
