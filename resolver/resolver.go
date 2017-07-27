package resolver

import "github.com/hzlmn/graph"

// getIndex helper for getting index of node in the list
func getIndex(node *graph.Node, list []*graph.Node) int {
	for pos, elem := range list {
		if node == elem {
			return pos
		}
	}

	return -1
}

// InArray checks if certain value is in list of items
func inArray(value *graph.Node, list []*graph.Node) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}

	return false
}

// walk through node connections and populate resolved list
// TODO: This code is very hacky should be replaced later
func walk(node *graph.Node, resolved *[]*graph.Node, unresolved *[]*graph.Node) {
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
func Resolve(startNode *graph.Node) []*graph.Node {
	var resolved []*graph.Node
	var unresolved []*graph.Node

	walk(startNode, &resolved, &unresolved)

	return resolved
}
