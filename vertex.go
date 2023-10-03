package dag

import (
	"fmt"
	"reflect"
)

// Vertex is a vertex in a graph.
type Vertex[T any] struct {
	// Name is the name of the Vertex.
	// It is used to compute the hashcode of the Vertex.
	Name string
	// Value is the value of the Vertex.
	// It can be of any type.
	Value T
}

// NewVertex returns a new Vertex instance.
func NewVertex[T any](name string, value T) Vertex[T] {
	return Vertex[T]{Name: name, Value: value}
}

// Hashcode returns the hashcode of the Vertex.
// It is computed using the Name and the type of the Value.
// It is intended to be used as a unique identifier for the Vertex in a graph.
func (v *Vertex[T]) hashcode() string {
	vType := reflect.TypeOf(v.Value)

	return fmt.Sprintf("%s-%s", v.Name, vType.Name())
}

// String returns the string representation of the Vertex.
func (v *Vertex[T]) String() string {
	return v.Name
}

// VertexID returns  the id of the Vertex.
func VertexID[T any](v Vertex[T]) string {
	return v.hashcode()
}
