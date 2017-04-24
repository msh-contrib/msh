package graph

import (
	"testing"
)

func TestGraphCreating(t *testing.T) {
	graph := New()
	if graph == nil {
		t.Log("Should create new graph instance")
		t.Fatal()
	}
}

func TestAddingNode(t *testing.T) {
	graph := New()
	graph.AddNode("testNode", make(map[string]Any))
	if graph.nodes["testNode"] != nil {
		t.Log("Not add node to graph")
		t.Fatal()
		if graph.nodes["testNode"].id != "testNode" {
			t.Log("Bad node")
			t.Fatal()
		}
	}
}

func TestAttachingNode(t *testing.T) {
	graph := New()
	testNode := NewNode("testNode", make(map[string]Any))
	graph.AttachNode(testNode)
	if graph.nodes["testNode"] != testNode {
		t.Log("Not correctly attaching node to graph")
		t.Fatal()
	}
}

//func TestGetNodes(t *testing.T) {
//graph := New()
//graph.AddNode("testNode", make)
//}
