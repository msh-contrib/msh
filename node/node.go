package node

type Node struct {
	params map[string]interface{}
	id     string
	edges  []*Node
}

func New(id string, params map[string]interface{}) *Node {
	return &Node{
		id:     id,
		params: params,
	}
}

func (n *Node) GetParams() map[string]interface{} {
	return n.params
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
