package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/rltran-codex/aoc-2024-go/utils"
	structures "github.com/rltran-codex/aoc-2024-go/utils/structures"
)

type GardenGraph struct {
	Farm      [][]string
	FarmGraph structures.Graph
	Regions   []Region
	Visited   map[*structures.GraphNode]bool
}

type Region struct {
	PlantType string
	Nodes     []*structures.GraphNode
	Area      int
	Perimeter int
	Sides     int
}

// post order traversal
func (g *GardenGraph) identifyRegion(currFlower *structures.GraphNode) (int, int, error) {
	if g.Visited[currFlower] {
		return 0, 0, fmt.Errorf("already processed")
	}

	// new region
	region := Region{
		PlantType: currFlower.Value.(string),
		Nodes:     []*structures.GraphNode{},
		Area:      0,
		Perimeter: 0,
	}

	// create a queue for BFS
	queue := []*structures.GraphNode{}
	queue = append(queue, currFlower)

	for len(queue) > 0 {
		curr := utils.PopQueue(&queue)
		if g.Visited[curr] {
			continue
		}

		// mark current node as visited and increment area
		g.Visited[curr] = true
		region.Nodes = append(region.Nodes, curr)
		region.Area++

		// queue non visited adjacent nodes
		for _, neighbor := range curr.Adj {
			if neighbor.Value != currFlower.Value {
				region.Perimeter++
			} else if !g.Visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}

		region.Perimeter += 4 - len(curr.Adj)
	}

	// add region to garden
	g.Regions = append(g.Regions, region)
	g.identifySides(&region)
	return region.Perimeter, region.Perimeter, nil
}

func (g *GardenGraph) identifySides(r *Region) {
	if len(r.Nodes) < 2 {
		r.Sides = 4
		return
	}
}

func (g *GardenGraph) isInbounds(x int, y int) bool {
	maxRow := len(g.Farm)
	maxCol := len(g.Farm[0])
	return x >= 0 && x < maxRow && y >= 0 && y < maxCol
}

// Sort nodes clockwise based on the angle relative to the centroid
func sortClockwise(nodes []*structures.GraphNode) {
	calculateAngle := func(x, y, centerX, centerY float64) float64 {
		return math.Atan2(y-centerY, x-centerX)
	}
	var centerX, centerY float64
	for _, node := range nodes {
		centerX += float64(node.X)
		centerY += float64(node.Y)
	}
	centerX /= float64(len(nodes))
	centerY /= float64(len(nodes))

	sort.Slice(nodes, func(i, j int) bool {
		// Calculate the angle for both nodes
		angleJ := calculateAngle(float64(nodes[j].X), float64(nodes[j].Y), centerX, centerY)
		angleI := calculateAngle(float64(nodes[i].X), float64(nodes[i].Y), centerX, centerY)
		return angleI < angleJ
	})
}

func main() {
	// main area to display puzzle answers
	garden := ParsePuzzleInput(false)
	fmt.Printf("Part 1: %d\n", Part1(garden))
	garden.Visited = make(map[*structures.GraphNode]bool)
	fmt.Printf("Part 2: %d\n", Part2(garden))
}

func ParsePuzzleInput(sample bool) *GardenGraph {
	farm := utils.Get2DPuzzleInput("day12.txt", sample)
	maxRow := len(farm)
	maxCol := len(farm[0])

	ngraph := structures.Graph{
		Nodes: make(map[string]*structures.GraphNode),
		Size:  0,
	}

	for i := range farm {
		for j := range farm[i] {
			key := utils.ParseKey(i, j)

			currNode, err := ngraph.GetGNode(key)
			if err != nil {
				currNode = &structures.GraphNode{
					Id:    key,
					Value: farm[i][j],
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
				key := utils.ParseKey(p.Row, p.Col)
				adjNode, err := ngraph.GetGNode(key)
				if err != nil {
					adjNode = &structures.GraphNode{
						Id:    key,
						Value: farm[p.Row][p.Col],
						X:     p.Row,
						Y:     p.Col,
					}
				}
				// queue node to be added
				adjNodes = append(adjNodes, adjNode)
			}

			// link graph node
			ngraph.AddGNode(currNode, adjNodes...)
		}
	}

	return &GardenGraph{
		FarmGraph: ngraph,
		Farm:      farm,
		Visited:   make(map[*structures.GraphNode]bool),
	}
}

func Part1(g *GardenGraph) int {
	result := 0
	for _, n := range g.FarmGraph.Nodes {
		area, perimeter, _ := g.identifyRegion(n)
		result += area * perimeter
	}

	return result
}
func Part2(g *GardenGraph) int {
	panic("Not yet implemented")
}
