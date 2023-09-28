package dag

import (
	"fmt"
)

// TopoOrder is a topological ordering.
type TopoOrder []Vertex

// TopologicalSort implements the kahn algorithm and returns a topological ordering of the graph.
// The graph must not have any cycle or it will return an error.
func (g *Graph) TopologicalSort() (*TopoOrder, error) {
	if len(g.adjacencyMap) == 0 {
		return &TopoOrder{}, nil
	}

	g2 := g.Copy()
	indegreeMap := make(map[Vertex]int)
	processIndegrees(indegreeMap, g)

	topOrder, visited := generateTopo(g2, indegreeMap)

	if visited != len(g.adjacencyMap) || visited == -1 {
		scc := g2.Cycles()
		return nil, fmt.Errorf("there exists a cycle in the graph: \n %#v", scc)
	}

	return topOrder, nil

}

func initQueue(graph *Graph, indegreeMap map[Vertex]int) TopoOrder {
	queue := make(TopoOrder, 0, len(graph.adjacencyMap))
	for hash := range graph.adjacencyMap {
		v := graph.hashMap[hash]
		if indegreeMap[v] == 0 {
			queue = append(queue, v)
		}
	}
	return queue
}

func generateTopo(graph *Graph, indegreeMap map[Vertex]int) (*TopoOrder, int) {
	topOrder := make(TopoOrder, 0, len(graph.adjacencyMap))
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
		for targetHash := range graph.adjacencyMap[hashcode(source)] {
			target := graph.hashMap[targetHash]
			indegreeMap[target]--
			if indegreeMap[target] == 0 {
				queue = append(queue, target)
			}
		}
	}
	return &topOrder, visited
}

func dequeue(queue *TopoOrder) {
	if len(*queue) > 1 {
		*queue = (*queue)[1:]
	} else {
		*queue = nil
	}
}

func processIndegrees(indegreeMap map[Vertex]int, graph *Graph) {
	/*
			for each source in Vertices
		    indegree[source] = 0
			for each edge(src, dest) in Edges
				indegree[dest]++
	*/
	for sourceHash, edges := range graph.adjacencyMap {
		source := graph.hashMap[sourceHash]
		if _, ok := indegreeMap[source]; !ok {
			indegreeMap[source] = 0
		}
		for targetHash := range edges {
			target := graph.hashMap[targetHash]
			indegreeMap[target]++
		}
	}

}
