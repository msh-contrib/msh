package graph

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/hzlmn/msh/node"
	"github.com/hzlmn/msh/utils"
	"github.com/mvdan/sh/syntax"
	//"reflect"
	//"io"
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
)

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

	g.nodes[nodeKey] = node.New(nodeKey, nodeParams)
}

// AttachNode add existed node with it dependencies to graph
func (g *Graph) AttachNode(node *node.Node) {
	for key := range g.nodes {
		if key == node.GetId() {
			return
		}
	}

	g.nodes[node.GetId()] = node

	for _, childNode := range node.GetConnections() {
		g.AttachNode(childNode)
	}
}

// GetNodes allow to get list of registered nodes in graph
func (g *Graph) GetNodes() (result []*node.Node) {
	for _, node := range g.nodes {
		result = append(result, node)
	}
	return
}

func CollectNodes(entry string) *Graph {
	graph := &Graph{}

	entryData := utils.ReadFile(entry)
	fileReader := bytes.NewReader(entryData)
	parsedFile, parseError := syntax.Parse(fileReader, "shell", syntax.PosixConformant)

	if parseError != nil {
		log.Fatal(parseError)
	}

	syntax.Walk(parsedFile, func(node syntax.Node) bool {
		switch x := node.(type) {
		case *syntax.Lit:
			fmt.Println("found word", spew.Sdump(x))
			break
		}

		return true
	})

	return graph
}
