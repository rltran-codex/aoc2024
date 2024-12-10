package datastructures

import (
	"fmt"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type GraphNode struct {
	Id    string
	Value interface{}
	Adj   []*GraphNode
}

// Only use if ID's will be unique
type Graph struct {
	Nodes map[string]*GraphNode
	Size  int
}

func addNode(pNode *GraphNode, cNode ...*GraphNode) {
	for _, c := range cNode {
		pNode.Adj = append(pNode.Adj, c)
		c.Adj = append(c.Adj, pNode)
	}
}

func connectionExists(pNode *GraphNode, cNode *GraphNode) bool {
	for _, v := range pNode.Adj {
		if cNode == v {
			return true
		}
	}

	return false
}

func removeNode(targetNode *GraphNode) {
	for i := 0; i < len(targetNode.Adj); i++ {
		n := targetNode.Adj[i]
		idx := utils.Index(n.Adj, targetNode)
		if idx == -1 {
			panic(fmt.Sprintf("%+v linked to %+v, but not found in %+v", targetNode, n, n))
		}

		n.Adj = append(n.Adj[:idx], n.Adj[idx+1:]...)
	}

	// clear fields to signify this is removed
	targetNode.Id = ""
	targetNode.Adj = nil
	targetNode.Value = nil
}

func (g *Graph) GetGNode(key string) (*GraphNode, error) {
	node, ok := g.Nodes[key]
	if !ok {
		return &GraphNode{}, fmt.Errorf("Graph node did not exist for '%s'", key)
	}

	return node, nil
}

// Function links the two nodes together and add them to a graph
func (g *Graph) AddGNode(pNode *GraphNode, cNode ...*GraphNode) {
	// Add parent to graph if DNE
	_, ok := g.Nodes[pNode.Id]
	if !ok {
		g.Nodes[pNode.Id] = pNode
	}

	for _, c := range cNode {
		if !connectionExists(pNode, c) {
			addNode(pNode, c)
		}

		// Add child to graph if DNE
		_, ok := g.Nodes[c.Id]
		if !ok {
			g.Nodes[c.Id] = c
		}
	}

	// update the size
	g.Size = len(g.Nodes)
}

// Function links removes a node from from the graph
func (g *Graph) RemoveGNode(targetNode *GraphNode) {
	removeNode(targetNode)

	// remove from the graph
	_, ok := g.Nodes[targetNode.Id]
	if !ok {
		panic(fmt.Sprintf("%+v was not found in Graph.Nodes", targetNode))
	}

	delete(g.Nodes, targetNode.Id)
	g.Size = len(g.Nodes)
}

func (g *Graph) DFSGraphTrav() {

}

func dfsRecur() {

}

func bfsTrav(g *Graph, startingNode *GraphNode, stop any) {
	panic("Unimplemented")
}
