package dag

// tarjan algorithm to detect cycles in a graph
// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
type tarjan[T any, V Vertex[T]] struct {
	graph    *Graph[T]
	index    int
	stack    []V
	indices  map[any]int
	lowlinks map[any]int
	onStack  map[any]bool
}

// Cycles detect and return cycles in the graph
// It implements the tarjan algorithm
// It returns a list of vertices that form a cycle
func Cycles[T any](g *Graph[T]) []Vertex[T] {
	if len(g.adjacencyMap) == 0 {
		return nil
	}

	var scc []Vertex[T]
	g2 := g.Copy()
	index := 0
	stack := make([]Vertex[T], 0, len(g2.adjacencyMap))
	tarjan := tarjan[T, Vertex[T]]{
		graph:    g2,
		index:    index,
		stack:    stack,
		indices:  make(map[any]int),
		lowlinks: make(map[any]int),
		onStack:  make(map[any]bool),
	}

	for _, v := range g2.hashMap {
		if _, ok := tarjan.indices[v.hashcode()]; !ok {
			scc = strongConnect(&v, &tarjan)
		}
	}

	if len(scc) > 0 {
		return scc
	}

	return nil

}

func strongConnect[T any](source *Vertex[T], tarjan *tarjan[T, Vertex[T]]) []Vertex[T] {
	tarjan.indices[source.hashcode()] = tarjan.index
	tarjan.lowlinks[source.hashcode()] = tarjan.index
	tarjan.index++
	tarjan.stack = append(tarjan.stack, *source)
	tarjan.onStack[source.hashcode()] = true
	output := []Vertex[T]{}

	src := source.hashcode()

	for edge := range tarjan.graph.adjacencyMap[source.hashcode()] {
		target := tarjan.graph.hashMap[edge]
		tar := edge
		if _, ok := tarjan.indices[tar]; !ok {
			strongConnect(&target, tarjan)
			tarjan.lowlinks[src] = min(tarjan.lowlinks[src], tarjan.lowlinks[tar])
		} else if tarjan.onStack[edge] {
			tarjan.lowlinks[src] = min(tarjan.lowlinks[src], tarjan.indices[tar])
		}
	}

	if tarjan.lowlinks[src] == tarjan.indices[src] {
		var w any
		var v Vertex[T]
		for w != src {
			v, tarjan.stack = tarjan.stack[len(tarjan.stack)-1], tarjan.stack[:len(tarjan.stack)-1]
			tarjan.onStack[src] = false
			output = append(output, v)
			w = v.hashcode()
		}
	}

	return output

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
