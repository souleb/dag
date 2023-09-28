package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyGraph(t *testing.T) {
	cases := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			"not empty int vertices",
			[]any{1, 2, 3},
			"\n1\n2\n3\n",
		},
		{
			"not empty string vertices",
			[]any{"a", "b", "c"},
			"\na\nb\nc\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}
			assert.Equal(t, g.String(), tc.expected)
		})

	}
}

func TestBasicGraph(t *testing.T) {
	cases := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			"not empty int vertices",
			[]any{1, 2, 3},
			"\n1\n  2 (1)\n  3 (1)\n2\n3\n",
		},
		{
			"not empty string vertices",
			[]any{"a", "b", "c"},
			"\na\n  b (1)\n  c (1)\nb\nc\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[2], 1)
			g.AddEdge(tc.input[0], tc.input[1], 1)

			assert.Equal(t, g.String(), tc.expected)
		})

	}
}

func TestGraphRemove(t *testing.T) {
	cases := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			"remove int vertex",
			[]any{1, 2, 3},
			"\n1\n  3 (1)\n3\n",
		},
		{
			"remove string vertex",
			[]any{"a", "b", "c"},
			"\na\n  c (1)\nc\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[2], 1)
			g.AddEdge(tc.input[0], tc.input[1], 1)
			g.Remove(tc.input[1])

			assert.Equal(t, g.String(), tc.expected)
		})

	}
}

func TestGraphRemoveEdge(t *testing.T) {
	cases := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			"remove int edge",
			[]any{1, 2, 3},
			"\n1\n  3 (1)\n2\n3\n",
		},
		{
			"remove string edge",
			[]any{"a", "b", "c"},
			"\na\n  c (1)\nb\nc\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[2], 1)
			g.AddEdge(tc.input[0], tc.input[1], 1)
			g.RemoveEdge(tc.input[0], tc.input[1])

			assert.Equal(t, g.String(), tc.expected)
		})

	}
}

func TestGraphHasEdge(t *testing.T) {
	cases := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			"has string edge",
			[]any{"a", "b", "c"},
			"a (b)",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}

			g.AddEdge(tc.input[0], tc.input[2], 1)
			g.AddEdge(tc.input[0], tc.input[1], 1)

			assert.True(t, g.HasEdge(tc.input[0], tc.input[1]))
		})

	}
}

func TestGraphHasVertex(t *testing.T) {
	cases := []struct {
		name     string
		input    []any
		expected string
	}{
		{
			"has string edge",
			[]any{"a", "b", "c"},
			"a",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}

			assert.True(t, g.HasVertex(tc.input[0]))
		})

	}
}
