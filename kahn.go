package dag

import (
	"fmt"
)

// TopoOrder is a topological ordering.
type TopoOrder[T any] []Vertex[T]

// TopologicalSort implements the kahn algorithm and returns a topological ordering of the graph.
// The graph must not have any cycle or it will return an error.
func (g *Graph[T]) TopologicalSort() (*TopoOrder[T], error) {
	if len(g.adjacencyMap) == 0 {
		return &TopoOrder[T]{}, nil
	}

	g2 := g.Copy()
	indegreeMap := make(map[string]int)
	processIndegrees(indegreeMap, g)

	topOrder, visited := generateTopo(g2, indegreeMap)

	if visited != len(g.adjacencyMap) || visited == -1 {
		scc := Cycles(g)
		return nil, fmt.Errorf("there exists a cycle in the graph: \n %#v", scc)
	}

	return topOrder, nil

}

func initQueue[T any, V Vertex[T]](graph *Graph[T], indegreeMap map[string]int) TopoOrder[T] {
	queue := make(TopoOrder[T], 0, len(graph.adjacencyMap))
	for hash := range graph.adjacencyMap {
		if indegreeMap[hash] == 0 {
			queue = append(queue, graph.hashMap[hash])
		}
	}
	return queue
}

func generateTopo[T any, V Vertex[T]](graph *Graph[T], indegreeMap map[string]int) (*TopoOrder[T], int) {
	topOrder := make(TopoOrder[T], 0, len(graph.adjacencyMap))
	var visited int

	queue := initQueue(graph, indegreeMap)
	if len(queue) == 0 {
		return nil, -1
	}

	for queue != nil {
		source := queue[0]
		dequeue(&queue)

		topOrder = append(topOrder, source)
		visited++
		for targetHash := range graph.adjacencyMap[source.hashcode()] {
			indegreeMap[targetHash]--
			if indegreeMap[targetHash] == 0 {
				queue = append(queue, graph.hashMap[targetHash])
			}
		}
	}
	return &topOrder, visited
}

func dequeue[T any](queue *TopoOrder[T]) {
	if len(*queue) > 1 {
		*queue = (*queue)[1:]
	} else {
		*queue = nil
	}
}

func processIndegrees[T any, V Vertex[T]](indegreeMap map[string]int, graph *Graph[T]) {
	/*
			for each source in Vertices
		    indegree[source] = 0
			for each edge(src, dest) in Edges
				indegree[dest]++
	*/
	for sourceHash, edges := range graph.adjacencyMap {
		if _, ok := indegreeMap[sourceHash]; !ok {
			indegreeMap[sourceHash] = 0
		}
		for targetHash := range edges {
			indegreeMap[targetHash]++
		}
	}

}
