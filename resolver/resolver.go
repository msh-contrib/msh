package resolver

import "github.com/hzlmn/msh/node"

func getIndex(node *node.Node, list []*node.Node) int {
	for pos, elem := range list {
		if node == elem {
			return pos
		}
	}

	return -1
}

// InArray checks if certain value is in list of items
func inArray(value *node.Node, list []*node.Node) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

func walk(node *node.Node, resolved *[]*node.Node, unresolved *[]*node.Node) {
	*unresolved = append(*unresolved, node)

	for _, edge := range node.GetConnections() {
		if ok := inArray(edge, *resolved); !ok {
			if ok := inArray(edge, *unresolved); ok {
				return
			}
			walk(edge, resolved, unresolved)
		}
	}

	*resolved = append(*resolved, node)
	index := getIndex(node, *unresolved)
	*unresolved = append((*unresolved)[:index], (*unresolved)[index+1:]...)
}

// Resolve run resoluting process starting from entry node
func Resolve(startNode *node.Node) []*node.Node {
	var resolved []*node.Node
	var unresolved []*node.Node

	walk(startNode, &resolved, &unresolved)

	return resolved
}
