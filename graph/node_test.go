package graph

import (
  "testing"
)

var node *Node

func TestNodeCreating(t *testing.T) {
  node = NewNode("testNode", make(map[string]interface{}))
  if node.id != "testNode" {
    t.Log("Uncorrect node identifier")
    t.Fatal()
  }
}

func TestNodeAssign(t *testing.T) {
  childNode := NewNode("childNode", make(map[string]interface{}))
  node.AddEdge(childNode)
  if node.edges[0] != childNode {
    t.Log("Not assign edges")
    t.Fatal()
  }
}

func TestGetConnections(t *testing.T) {
  nodeList := node.GetConnections()
  if len(nodeList) < 1 {
    t.Log("Not return list of edges")
    t.Fatal()
  }
}
