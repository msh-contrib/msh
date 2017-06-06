package graph

import (
	//"github.com/davecgh/go-spew/spew"
	"github.com/hzlmn/msh/node"
	"github.com/hzlmn/msh/utils"
	//"reflect"
	//"io"
	//"fmt"
	"log"
	"os"
	//"path/filepath"
	"path/filepath"
	"regexp"
)

// Graph defines graph structure
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
func (g *Graph) AddNode(nodeKey string, nodeParams map[string]interface{}) {
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

// AddEdge adds links between two nodes by id
func (g *Graph) AddEdge(startID string, endID string) {
	g.AddNode(startID, make(map[string]interface{}))
	g.AddNode(endID, make(map[string]interface{}))
	g.nodes[startID].AddEdge(g.nodes[endID])
}

// GetNode find node by id
func (g *Graph) GetNode(nodeID string) *node.Node {
	return g.nodes[nodeID]
}

// NormalizePath normalize file path
func NormalizePath(fpath string) string {
	currDir, _ := os.Getwd()

	if filepath.IsAbs(fpath) {
		return fpath
	}

	return filepath.Join(currDir, fpath)
}

// Ctx defines walk function context of execution
type Ctx struct {
	Graph    *Graph
	Reg      *regexp.Regexp
	Dir      string
	CurrFile string
	ParentID string
}

// walk recursively walks files and creates deps graph
// based on include statements
func walk(ctx *Ctx) {
	normPath := NormalizePath(ctx.CurrFile)

	if ctx.Graph.GetNode(normPath) != nil {
		// Ignore if node already presented
		return
	}

	fdata := string(utils.ReadFile(normPath))

	ctx.Graph.AddNode(normPath, map[string]interface{}{
		"data": fdata,
	})

	if ctx.ParentID != "" {
		ctx.Graph.AddEdge(ctx.ParentID, normPath)
	}

	//result := reqRegExp.FindAllStringSubmatch(entryData, -1)

	for _, match := range ctx.Reg.FindAllStringSubmatch(fdata, -1) {
		ctx.ParentID = normPath
		ctx.CurrFile = match[1]
		walk(ctx)
	}

}

// MakeGraph run dependency graph creating process starting from entry node
func MakeGraph(entry string) *Graph {
	graph := New()
	currDir, _ := os.Getwd()
	reqRegExp, regError := regexp.Compile("include\\s+'(.*?)'")

	if regError != nil {
		log.Fatal("Error while compiling RegExp", regError)
	}

	walk(&Ctx{
		Graph:    graph,
		Reg:      reqRegExp,
		Dir:      currDir,
		CurrFile: entry,
	})

	return graph
}
