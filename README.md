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
  // Create a new DAG
  d := dag.Graph{}

  // Add some nodes
  d.Add("a")
  d.Add("b")
  d.Add("c")
  d.Add("d")
  d.Add("e")
  d.Add("f")

  // Add some edges
  d.AddEdge("a", "b", 1)
  d.AddEdge("a", "c", 1)
  d.AddEdge("b", "d", 1)
  d.AddEdge("c", "d", 1)

  // Sort the DAG
  sorted, err := d.TopologicalSort()
  if err != nil {
    panic(err)
  }

  // Print the sorted nodes
  fmt.Println(sorted.String())
}
```