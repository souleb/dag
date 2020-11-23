package dag

import (
	"strings"
	"testing"
)

func TestEmptyGraph(t *testing.T) {
	cases := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			"not empty int vertices",
			[]interface{}{1, 2, 3},
			"1\n2\n3\n",
		},
		{
			"not empty string vertices",
			[]interface{}{"a", "b", "c"},
			"a\nb\nc\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}
			actual := strings.TrimSpace(g.String())
			expected := strings.TrimSpace(tc.expected)
			if actual != expected {
				t.Fatalf("Test %s \n expected: %s\n actual: %s", tc.name, expected, actual)
			}
		})

	}
}

func TestBasicGraph(t *testing.T) {
	cases := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			"not empty int vertices",
			[]interface{}{1, 2, 3},
			"1\n  2 (1)\n  3 (1)\n2\n3",
		},
		{
			"not empty string vertices",
			[]interface{}{"a", "b", "c"},
			"a\n  b (1)\n  c (1)\nb\nc",
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

			actual := strings.TrimSpace(g.String())
			expected := strings.TrimSpace(tc.expected)
			if actual != expected {
				t.Fatalf("Test %s \n expected:\n%s\n actual:\n%s", tc.name, expected, actual)
			}
		})

	}
}

func TestGraphRemove(t *testing.T) {
	cases := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			"remove int vertex",
			[]interface{}{1, 2, 3},
			"1\n  3 (1)\n3",
		},
		{
			"remove string vertex",
			[]interface{}{"a", "b", "c"},
			"a\n  c (1)\nc",
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

			actual := strings.TrimSpace(g.String())
			expected := strings.TrimSpace(tc.expected)
			if actual != expected {
				t.Fatalf("Test %s \n expected:\n%s\n actual:\n%s", tc.name, expected, actual)
			}
		})

	}
}

func TestGraphRemoveEdge(t *testing.T) {
	cases := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			"remove int edge",
			[]interface{}{1, 2, 3},
			"1\n  3 (1)\n2\n3",
		},
		{
			"remove string edge",
			[]interface{}{"a", "b", "c"},
			"a\n  c (1)\nb\nc",
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

			actual := strings.TrimSpace(g.String())
			expected := strings.TrimSpace(tc.expected)
			if actual != expected {
				t.Fatalf("Test %s \n expected:\n%s\n actual:\n%s", tc.name, expected, actual)
			}
		})

	}
}

func TestGraphHasEdge(t *testing.T) {
	cases := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			"has string edge",
			[]interface{}{"a", "b", "c"},
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

			if !g.HasEdge(tc.input[0], tc.input[1]) {
				t.Fatalf("Test %s \n should have %s", tc.name, tc.expected)
			}
		})

	}
}

func TestGraphHasVertex(t *testing.T) {
	cases := []struct {
		name     string
		input    []interface{}
		expected string
	}{
		{
			"has string edge",
			[]interface{}{"a", "b", "c"},
			"a",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
			for _, v := range tc.input {
				g.Add(v)
			}

			if !g.HasVertex(tc.input[0]) {
				t.Fatalf("Test %s \n should have %s", tc.name, tc.expected)
			}
		})

	}
}
