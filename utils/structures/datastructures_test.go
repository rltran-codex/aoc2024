package datastructures

import (
	"testing"

	"github.com/rltran-codex/aoc-2024-go/utils"
)

func TestNodeOp(t *testing.T) {
	node1 := GraphNode{Id: "1", Value: 1}
	node2 := GraphNode{Id: "2", Value: 2}
	node3 := GraphNode{Id: "3", Value: 3}
	node4 := GraphNode{Id: "4", Value: 4}

	AddNode(&node1, &node2)

	// test if node1 is prev for node 2
	if node1.Next[0] != &node2 {
		t.Errorf("Node2 was not found in Node1.Next")
	}
	// test if node2 is next for node 1
	if node2.Prev[0] != &node1 {
		t.Errorf("Node1 was not found in Node2.Prev")
	}

	AddNode(&node1, &node3)
	AddNode(&node1, &node4)
	AddNode(&node2, &node3)
	AddNode(&node2, &node4)
	RemoveNode(&node2)

	idx := utils.Index(node1.Next, &node2)
	if idx != -1 {
		t.Errorf("Node2 was found in Node1.Next: %+v", node1.Next)
	}
	idx = utils.Index(node3.Prev, &node2)
	if idx != -1 {
		t.Errorf("Node2 was found in Node3.Prev: %+v", node1.Next)
	}
	idx = utils.Index(node4.Prev, &node2)
	if idx != -1 {
		t.Errorf("Node2 was found in Node3.Prev: %+v", node1.Next)
	}
}
