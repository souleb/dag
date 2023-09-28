package dag

import "fmt"

// Vertex of the graph
// can be anything
type Vertex interface{}

// VertexHashable is an optional interface that can be implemented to specify
// an alternate hash code for a Vertex. If this isnt implemented, Go interface
// equality is used.
type VertexHashable interface {
	Hashcode() interface{}
}

// VertexID returns  the id of the Vertex.
func VertexID(v Vertex) interface{} {
	return hashcode(v)
}

// VertexName returns the nam eof a vertex
func VertexName(v Vertex) string {
	switch v := v.(type) {
	case fmt.Stringer:
		return fmt.Sprintf("%s", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// hashcode returns the hascode for a Vertex.
func hashcode(v interface{}) interface{} {
	if h, ok := v.(VertexHashable); ok {
		return h.Hashcode()
	}
	return v
}
