package dag

import (
	"fmt"
	"sort"
	"strings"
)

// EdgeSet is a set data structure
type EdgeSet map[any]int

// Graph represents a directed acyclic graph
type Graph struct {
	adjacencyMap map[any]EdgeSet
	// maintain a mapping of the vertices to their hashes.
	hashMap map[any]Vertex
}

// Add a Vertex to the adjacencyMap.
// It a Vertex with the same identity exist, it gets overwritten.
func (g *Graph) Add(v Vertex) {
	g.init()
	// Add the vertex entry
	hash := hashcode(v)
	if _, ok := g.adjacencyMap[hash]; !ok {
		g.adjacencyMap[hash] = make(EdgeSet)
	}

	g.hashMap[hash] = v
}

// Remove delete a Vertex from the adjacencyMap.
func (g *Graph) Remove(v Vertex) {
	hash := hashcode(v)
	// delete the vertex entry
	delete(g.adjacencyMap, hash)

	// delete all occurence of the Vertex in the sets.
	for _, set := range g.adjacencyMap {
		delete(set, hash)
	}

	delete(g.hashMap, hash)
}

// Vertex returns a Vertex pointer if it exists in the graph
// Othervise returns a false boolean
func (g *Graph) Vertex(hash string) (Vertex, bool) {
	if v, ok := g.hashMap[hash]; ok {
		return v, true
	}

	return nil, false
}

// AddEdge add an edge to the Graph.
func (g *Graph) AddEdge(source, target Vertex, weight int) {
	g.init()

	// Make sure that every used vertex shows up in our map keys.
	hashSource, hashTarget := hashcode(source), hashcode(target)
	if _, ok := g.adjacencyMap[hashSource]; !ok {
		g.Add(source)
	}

	if _, ok := g.adjacencyMap[hashTarget]; !ok {
		g.Add(target)
	}

	g.adjacencyMap[hashSource][hashTarget] = weight
}

// RemoveEdge delete an edge from the adjacencyMap.
func (g *Graph) RemoveEdge(source, target Vertex) {
	hashSource, hashTarget := hashcode(source), hashcode(target)
	if set, ok := g.adjacencyMap[hashSource]; ok {
		delete(set, hashTarget)
	}
}

// HasEdge check if an edge exist between to vertices.
func (g *Graph) HasEdge(source, target Vertex) bool {
	hashSource, hashTarget := hashcode(source), hashcode(target)
	if set, ok := g.adjacencyMap[hashSource]; ok {
		if _, ok := set[hashTarget]; ok {
			return true
		}
	}
	return false
}

// HasVertex check if a vertex is in the adjacencyMap.
func (g *Graph) HasVertex(v Vertex) bool {
	hash := hashcode(v)
	if _, ok := g.adjacencyMap[hash]; ok {
		return true
	}
	return false
}

// computeGraphRepresentation returns a slice of vertices and a map of edges
// for a graph.
func (g *Graph) computeGraphRepresentation() ([]string, map[string][]string) {
	names := make([]string, 0, len(g.adjacencyMap))
	mapping := make(map[string][]string, len(g.adjacencyMap))

	// Get the vertices and edges in alphabetical orders by using a string sort.
	// having this deterministic behavior make testing easier.
	for v, targets := range g.adjacencyMap {
		names = append(names, VertexName(v))
		deps := make([]string, 0, len(targets))

		for target, weight := range targets {
			deps = append(deps, fmt.Sprintf(
				"%s (%d)", VertexName(target), weight))
		}
		sort.Strings(deps)

		mapping[VertexName(v)] = deps

	}

	sort.Strings(names)

	return names, mapping
}

// Copy returns a copy of the graph
func (g *Graph) Copy() *Graph {
	var newGraph Graph
	newGraph.init()

	for hash, edges := range g.adjacencyMap {
		newEdges := make(EdgeSet)
		for edge, weight := range edges {
			newEdges[edge] = weight
		}
		newGraph.adjacencyMap[hash] = newEdges
	}

	for hash, vertex := range g.hashMap {
		newGraph.hashMap[hash] = vertex
	}

	return &newGraph
}

// String is a human-friendly representation of the graph
func (g *Graph) String() string {
	var buf strings.Builder
	buf.WriteString("\n")

	names, mapping := g.computeGraphRepresentation()

	for _, name := range names {
		buf.WriteString(fmt.Sprintf("%s\n", name))
		for _, d := range mapping[name] {
			buf.WriteString(fmt.Sprintf("  %s\n", d))
		}
	}

	return buf.String()
}

func (g *Graph) init() {
	if g.adjacencyMap == nil {
		g.adjacencyMap = make(map[any]EdgeSet)
	}

	if g.hashMap == nil {
		g.hashMap = make(map[any]Vertex)
	}
}
