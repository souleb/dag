package dag

// tarjan algorithm to detect cycles in a graph
// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm
type tarjan struct {
	graph    *Graph
	index    int
	stack    []Vertex
	indices  map[any]int
	lowlinks map[any]int
	onStack  map[any]bool
}

// Cycles detect and return cycles in the graph
// It implements the tarjan algorithm
// It returns a list of vertices that form a cycle
func (g *Graph) Cycles() []Vertex {
	if len(g.adjacencyMap) == 0 {
		return nil
	}

	var scc []Vertex
	g2 := g.Copy()
	index := 0
	stack := make([]Vertex, 0, len(g2.adjacencyMap))
	tarjan := tarjan{
		graph:    g2,
		index:    index,
		stack:    stack,
		indices:  make(map[any]int),
		lowlinks: make(map[any]int),
		onStack:  make(map[any]bool),
	}

	for _, v := range g2.hashMap {
		if _, ok := tarjan.indices[hashcode(v)]; !ok {
			scc = strongConnect(&v, &tarjan)
		}
	}

	if len(scc) > 0 {
		return scc
	}

	return nil

}

func strongConnect(source *Vertex, tarjan *tarjan) []Vertex {
	tarjan.indices[hashcode(*source)] = tarjan.index
	tarjan.lowlinks[hashcode(*source)] = tarjan.index
	tarjan.index++
	tarjan.stack = append(tarjan.stack, *source)
	tarjan.onStack[hashcode(*source)] = true
	output := []Vertex{}

	src := hashcode(*source)

	for edge := range tarjan.graph.adjacencyMap[hashcode(*source)] {
		target := tarjan.graph.hashMap[hashcode(edge)]
		tar := hashcode(edge)
		if _, ok := tarjan.indices[tar]; !ok {
			strongConnect(&target, tarjan)
			tarjan.lowlinks[src] = min(tarjan.lowlinks[src], tarjan.lowlinks[tar])
		} else if tarjan.onStack[hashcode(tar)] {
			tarjan.lowlinks[src] = min(tarjan.lowlinks[src], tarjan.indices[tar])
		}
	}

	if tarjan.lowlinks[src] == tarjan.indices[src] {
		var w any
		var v Vertex
		for w != src {
			v, tarjan.stack = tarjan.stack[len(tarjan.stack)-1], tarjan.stack[:len(tarjan.stack)-1]
			tarjan.onStack[src] = false
			output = append(output, v)
			w = hashcode(v)
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
