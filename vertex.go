package dag

import "fmt"

// Vertex of the graph
// can be anything
type Vertex any

// VertexHashable is an optional interface that can be implemented to specify
// an alternate hash code for a Vertex.
type VertexHashable interface {
	Hashcode() any
}

// VertexID returns  the id of the Vertex.
func VertexID(v Vertex) any {
	return hashcode(v)
}

// VertexName returns the name of a vertex
func VertexName(v Vertex) string {
	switch v := v.(type) {
	case fmt.Stringer:
		return v.String()
	default:
		return fmt.Sprintf("%v", v)
	}
}

// hashcode returns the hascode for a Vertex.
func hashcode(v any) any {
	if h, ok := v.(VertexHashable); ok {
		return h.Hashcode()
	}
	return v
}
