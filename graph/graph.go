package graph

type Graph struct {
  nodes map[string]*Node
}

func New() *Graph {
  return &Graph{
    nodes: make(map[string]*Node),
  }
}

// Create and add new node to graph
func (g *Graph) AddNode(nodeKey string, nodeParams interface{}) {
  for key := range g.nodes {
    if key == nodeKey {
      return
    }
  }

  g.nodes[nodeKey] = NewNode(nodeKey, nodeParams)
}

// Add existed node with it dependencies to graph
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

// Get list of registered nodes in graph
func (g *Graph) GetNodes() (result []*Node) {
  for _, node := range g.nodes {
    result = append(result, node)
  }
  return
}
