package main

import (
  "fmt"
  "github.com/hzlmn/msh/graph"
)

func main() {
  params := map[string]interface{}{
    "param": "test",
  }

  childNode := graph.NewNode(params, "child")
  node := graph.NewNode(params, "key")

  node.AddEdge(childNode)

  graph := graph.New()
  graph.AttachNode(node)

  fmt.Println(graph.GetNodes())
}
