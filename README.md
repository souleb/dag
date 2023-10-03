# dag

Dag is a simple, lightweight, and easy to use directed acyclic graph (DAG) library for Go.

It provides a simple API for creating and manipulating DAGs, and a topological sort function for sorting the DAG.

```go
package main

import (
  "fmt"

  "github.com/souleb/dag"
)

func main() {
  // Create a new DAG with vertices of type string
  d := dag.New[string]

  // Create some vertices
  a := dag.NewVertex[string]("a")
  b := dag.NewVertex[string]("b")
  c := dag.NewVertex[string]("c")
  d := dag.NewVertex[string]("d")
  e := dag.NewVertex[string]("e")

// Add the vertices to the DAG
  d.AddVertex(a)
  d.AddVertex(b)
  d.AddVertex(c)
  d.AddVertex(d)
  d.AddVertex(e)

  // Add some edges
  d.AddEdge(a, b, 1)
  d.AddEdge(a, c, 1)
  d.AddEdge(b, d, 1)
  d.AddEdge(c, d, 1)
  d.AddEdge(d, e, 1)

  // Sort the DAG
  sorted, err := d.TopologicalSort()
  if err != nil {
    panic(err)
  }

  // Print the sorted nodes
  fmt.Println(sorted.String())
}
```