package graph

type Any interface{}

type Node struct {
	params Any
	id     string
	edges  []*Node
}

func NewNode(id string, params Any) *Node {
	return &Node{
		id:     id,
		params: params,
	}
}

func (n *Node) AddEdge(node *Node) {
	for _, item := range n.edges {
		if item == node {
			return
		}
	}

	n.edges = append(n.edges, node)
}

func (n *Node) GetConnections() []*Node {
	return n.edges
}

func (n *Node) GetId() string {
	return n.id
}
