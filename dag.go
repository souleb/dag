package dag

import (
	"fmt"
	"sort"
	"strings"
)

// EdgeSet is a set data structure for edges.
type EdgeSet map[string]int

// Graph represents a directed acyclic graph
type Graph[T any] struct {
	adjacencyMap map[string]EdgeSet
	// maintain a mapping of the vertices to their hashes.
	hashMap map[string]Vertex[T]
}

// New returns a new Graph instance.
func New[T any]() *Graph[T] {
	g := Graph[T]{}
	g.init()
	return &g
}

// Add a Vertex to the adjacencyMap.
// It a Vertex with the same identity exist, it gets overwritten.
func (g *Graph[T]) Add(v Vertex[T]) {
	g.init()
	// Add the vertex entry
	hash := v.hashcode()
	if _, ok := g.adjacencyMap[hash]; !ok {
		g.adjacencyMap[hash] = make(EdgeSet)
	}

	g.hashMap[hash] = v
}

// Remove delete a Vertex from the adjacencyMap.
func (g *Graph[T]) Remove(v Vertex[T]) {
	hash := v.hashcode()
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
func (g *Graph[T]) Vertex(hash string) (Vertex[T], bool) {
	if v, ok := g.hashMap[hash]; ok {
		return v, true
	}

	var empty Vertex[T]
	return empty, false
}

// AddEdge add an edge to the Graph.
func (g *Graph[T]) AddEdge(source, target Vertex[T], weight int) {
	g.init()

	// Make sure that every used vertex shows up in our map keys.
	hashSource, hashTarget := source.hashcode(), target.hashcode()
	if _, ok := g.adjacencyMap[hashSource]; !ok {
		g.Add(source)
	}

	if _, ok := g.adjacencyMap[hashTarget]; !ok {
		g.Add(target)
	}

	g.adjacencyMap[hashSource][hashTarget] = weight
}

// RemoveEdge delete an edge from the adjacencyMap.
func (g *Graph[T]) RemoveEdge(source, target Vertex[T]) {
	hashSource, hashTarget := source.hashcode(), target.hashcode()
	if set, ok := g.adjacencyMap[hashSource]; ok {
		delete(set, hashTarget)
	}
}

// HasEdge check if an edge exist between to vertices.
func (g *Graph[T]) HasEdge(source, target Vertex[T]) bool {
	hashSource, hashTarget := source.hashcode(), target.hashcode()
	if set, ok := g.adjacencyMap[hashSource]; ok {
		if _, ok := set[hashTarget]; ok {
			return true
		}
	}
	return false
}

// HasVertex check if a vertex is in the adjacencyMap.
func (g *Graph[T]) HasVertex(v Vertex[T]) bool {
	hash := v.hashcode()
	if _, ok := g.adjacencyMap[hash]; ok {
		return true
	}
	return false
}

// computeGraphRepresentation returns a slice of vertices and a map of edges
// for a graph.
func (g *Graph[T]) computeGraphRepresentation() ([]string, map[string][]string) {
	names := make([]string, 0, len(g.adjacencyMap))
	mapping := make(map[string][]string, len(g.adjacencyMap))

	// Get the vertices and edges in alphabetical orders by using a string sort.
	// having this deterministic behavior make testing easier.
	for v, targets := range g.adjacencyMap {
		names = append(names, v)
		deps := make([]string, 0, len(targets))

		for target, weight := range targets {
			deps = append(deps, fmt.Sprintf(
				"%s (%d)", target, weight))
		}
		sort.Strings(deps)

		mapping[v] = deps

	}

	sort.Strings(names)

	return names, mapping
}

// Copy returns a copy of the graph
func (g *Graph[T]) Copy() *Graph[T] {
	var newGraph Graph[T]
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
func (g *Graph[T]) String() string {
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

func (g *Graph[T]) init() {
	if g.adjacencyMap == nil {
		g.adjacencyMap = make(map[string]EdgeSet)
	}

	if g.hashMap == nil {
		g.hashMap = make(map[string]Vertex[T])
	}
}
