package graph

import "github.com/hzlmn/msh/node"

type Graph struct {
	nodes map[string]*node.Node
}

// New create new instance of graph structure
func New() *Graph {
	return &Graph{
		nodes: make(map[string]*node.Node),
	}
}

// AddNode create and add new node to graph
func (g *Graph) AddNode(nodeKey string, nodeParams interface{}) {
	for key := range g.nodes {
		if key == nodeKey {
			return
		}
	}

	g.nodes[nodeKey] = node.NewNode(nodeKey, nodeParams)
}

// AttachNode add existed node with it dependencies to graph
func (g *Graph) AttachNode(node *Node) {
	for key := range g.nodes {
		if key == node.id {
			return
		}
	}

	g.nodes[node.id] = node

	for _, childNode := range node.edges {
		g.AttachNode(childNode)
	}
}

// GetNodes allow to get list of registered nodes in graph
func (g *Graph) GetNodes() (result []*Node) {
	for _, node := range g.nodes {
		result = append(result, node)
	}
	return
}
