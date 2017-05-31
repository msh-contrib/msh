package resolver

import (
	"fmt"
	"github.com/hzlmn/msh/node"
)

// Resolver represents
type Resolver struct {
	entryNode *node.Node
}

// New creates new resolver structure
func New(node *node.Node) *Resolver {
	return &Resolver{
		entryNode: node,
	}
}

// Resolve run resoluting process starting from entry node
func (r *Resolver) Resolve(startNode *node.Node) []*node.Node {
	var unresolved []*node.Node
	var resolved []*node.Node

	walk := func(node *node.Node, resolved []*node.Node, unresolved []*node.Node) {
		unresolved = append(unresolved, node)
		//edges := node.GetConnections()
		for _, edge := range node.GetConnections() {
			fmt.Println(edge)
		}
		return
	}

	walk(startNode, resolved, unresolved)

	return resolved
}
