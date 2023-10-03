package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyGraph(t *testing.T) {
	cases := []struct {
		name     string
		input    []Vertex[any]
		expected string
	}{
		{
			"not empty int vertices",
			[]Vertex[any]{{Name: "1", Value: 1}, {Name: "2", Value: 2}, {Name: "3", Value: 3}},
			"\n1-int\n2-int\n3-int\n",
		},
		{
			"not empty string vertices",
			[]Vertex[any]{{Name: "a", Value: "a"}, {Name: "b", Value: "b"}, {Name: "c", Value: "c"}},
			"\na-string\nb-string\nc-string\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph[any]{}
			for _, v := range tc.input {
				g.Add(v)
			}
			assert.Equal(t, tc.expected, g.String())
		})

	}
}

func TestBasicGraph(t *testing.T) {
	cases := []struct {
		name     string
		input    []Vertex[any]
		expected string
	}{
		{
			"not empty int vertices",
			[]Vertex[any]{{Name: "1", Value: 1}, {Name: "2", Value: 2}, {Name: "3", Value: 3}},
			"\n1-int\n  2-int (1)\n  3-int (1)\n2-int\n3-int\n",
		},
		{
			"not empty string vertices",
			[]Vertex[any]{{Name: "a", Value: "a"}, {Name: "b", Value: "b"}, {Name: "c", Value: "c"}},
			"\na-string\n  b-string (1)\n  c-string (1)\nb-string\nc-string\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph[any]{}
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[1], 1)
			g.AddEdge(tc.input[0], tc.input[2], 1)
			assert.Equal(t, g.String(), tc.expected)
		})

	}
}

func TestGraphRemove(t *testing.T) {
	cases := []struct {
		name     string
		input    []Vertex[any]
		expected string
	}{
		{
			"remove int vertex",
			[]Vertex[any]{{Name: "1", Value: 1}, {Name: "2", Value: 2}, {Name: "3", Value: 3}},
			"\n1-int\n  3-int (1)\n3-int\n",
		},
		{
			"remove string vertex",
			[]Vertex[any]{{Name: "a", Value: "a"}, {Name: "b", Value: "b"}, {Name: "c", Value: "c"}},
			"\na-string\n  c-string (1)\nc-string\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph[any]{}
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[1], 1)
			g.AddEdge(tc.input[0], tc.input[2], 1)
			g.Remove(tc.input[1])

			assert.Equal(t, tc.expected, g.String())
		})

	}
}

func TestGraphRemoveEdge(t *testing.T) {
	cases := []struct {
		name     string
		input    []Vertex[any]
		expected string
	}{
		{
			"remove int edge",
			[]Vertex[any]{{Name: "1", Value: 1}, {Name: "2", Value: 2}, {Name: "3", Value: 3}},
			"\n1-int\n  3-int (1)\n2-int\n3-int\n",
		},
		{
			"remove string edge",
			[]Vertex[any]{{Name: "a", Value: "a"}, {Name: "b", Value: "b"}, {Name: "c", Value: "c"}},
			"\na-string\n  c-string (1)\nb-string\nc-string\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph[any]{}
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[1], 1)
			g.AddEdge(tc.input[0], tc.input[2], 1)
			g.RemoveEdge(tc.input[0], tc.input[1])

			assert.Equal(t, tc.expected, g.String())
		})

	}
}

func TestGraphHasEdge(t *testing.T) {
	cases := []struct {
		name     string
		input    []Vertex[any]
		expected string
	}{
		{
			"has string edge",
			[]Vertex[any]{{Name: "a", Value: "a"}, {Name: "b", Value: "b"}, {Name: "c", Value: "c"}},
			"a (b)",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph[any]{}
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[1], 1)
			g.AddEdge(tc.input[0], tc.input[2], 1)

			assert.True(t, g.HasEdge(tc.input[0], tc.input[1]))
		})

	}
}

func TestGraphHasVertex(t *testing.T) {
	cases := []struct {
		name     string
		input    []Vertex[any]
		expected string
	}{
		{
			"has string edge",
			[]Vertex[any]{{Name: "a", Value: "a"}, {Name: "b", Value: "b"}, {Name: "c", Value: "c"}},
			"a",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph[any]{}
			for _, v := range tc.input {
				g.Add(v)
			}

			assert.True(t, g.HasVertex(tc.input[0]))
		})

	}
}
