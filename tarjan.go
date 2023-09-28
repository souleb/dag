package dag

type config struct {
	graph    *Graph
	index    *int
	stack    *[]Vertex
	indices  map[string]int
	lowlinks map[string]int
	onStack  map[string]bool
}

// Cycles detect and return cycles in the graph
// It implements the tarjan algorithm
// It returns a list of the first cycle found
func (g *Graph) Cycles() []Vertex {
	g2 := g.Copy()
	index := 0
	stack := make([]Vertex, 0, len(g2.adjacencyMap))
	conf := config{
		graph:    g2,
		index:    &index,
		stack:    &stack,
		indices:  make(map[string]int),
		lowlinks: make(map[string]int),
		onStack:  make(map[string]bool),
	}

	for _, v := range g2.hashMap {
		scc := strongConnect(&v, &conf)
		if len(scc) > 0 {
			return scc
		}
	}

	return nil

}

func strongConnect(source *Vertex, conf *config) []Vertex {
	conf.indices[hashcode(*source).(string)] = *conf.index
	conf.lowlinks[hashcode(*source).(string)] = *conf.index
	*conf.index++
	*conf.stack = append(*conf.stack, *source)
	conf.onStack[hashcode(*source).(string)] = true
	output := []Vertex{}

	src := hashcode(*source).(string)

	for edge := range conf.graph.adjacencyMap[hashcode(*source)] {
		target := conf.graph.hashMap[hashcode(edge)]
		tar := hashcode(edge).(string)
		if _, ok := conf.indices[tar]; !ok {
			strongConnect(&target, conf)
			conf.lowlinks[src] = min(conf.lowlinks[src], conf.lowlinks[tar])
		} else if conf.onStack[hashcode(tar).(string)] {
			conf.lowlinks[src] = min(conf.lowlinks[src], conf.indices[tar])
		}
	}

	if conf.lowlinks[src] == conf.indices[src] {
		var w string
		var v Vertex
		for w != src {
			v, *conf.stack = (*conf.stack)[len(*conf.stack)-1], (*conf.stack)[:len(*conf.stack)-1]
			output = append(output, v)
			w = hashcode(v).(string)
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
