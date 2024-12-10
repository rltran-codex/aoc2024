package datastructures

import (
	"fmt"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

type GraphNode struct {
	Id    string
	Value any
	Prev  []*GraphNode
	Next  []*GraphNode
}

// Only use if ID's will be unique
type Graph struct {
	Nodes map[string]*GraphNode
	Size  int
}

func AddNode(pNode *GraphNode, cNode ...*GraphNode) {
	for _, c := range cNode {
		pNode.Next = append(pNode.Next, c)
		c.Prev = append(c.Prev, pNode)
	}
}

func RemoveNode(targetNode *GraphNode) {
	for i := 0; i < len(targetNode.Next); i++ {
		n := targetNode.Next[i]
		idx := utils.Index(n.Prev, targetNode)
		if idx == -1 {
			panic(fmt.Sprintf("%+v linked to %+v, but not found in %+v", targetNode, n, n))
		}

		n.Prev = append(n.Prev[:idx], n.Prev[idx+1:]...)
	}
	for i := 0; i < len(targetNode.Prev); i++ {
		n := targetNode.Prev[i]
		idx := utils.Index(n.Next, targetNode)
		if idx == -1 {
			panic(fmt.Sprintf("%+v linked to %+v, but not found in %+v", targetNode, n, n))
		}

		n.Next = append(n.Next[:idx], n.Next[idx+1:]...)
	}

	// clear fields to signify this is removed
	targetNode.Id = ""
	targetNode.Next = nil
	targetNode.Prev = nil
	targetNode.Value = nil
}

// Function links the two nodes together and add them to a graph
func AddGNode(g *Graph, pNode *GraphNode, cNode ...*GraphNode) {
	// Add parent to graph if DNE
	_, ok := g.Nodes[pNode.Id]
	if !ok {
		g.Nodes[pNode.Id] = pNode
	}

	for _, c := range cNode {
		pNode.Next = append(pNode.Next, c)
		c.Prev = append(c.Prev, pNode)

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
func RemoveGNode(g *Graph, targetNode *GraphNode) {
	RemoveNode(targetNode)

	// remove from the graph
	_, ok := g.Nodes[targetNode.Id]
	if !ok {
		panic(fmt.Sprintf("%+v was not found in Graph.Nodes", targetNode))
	}

	delete(g.Nodes, targetNode.Id)
	g.Size = len(g.Nodes)
}

func DFSGraphTrav() {

}

func dfsRecur() {

}

func bfsTrav(g *Graph, startingNode *GraphNode, stop any) {
	panic("Unimplemented")
}
