package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTopologicalSort(t *testing.T) {
	cases := []struct {
		name     string
		input    map[Vertex[any]][]Vertex[any]
		expected *TopoOrder[any]
		wantErr  bool
	}{
		{
			"empty graph",
			map[Vertex[any]][]Vertex[any]{},
			&TopoOrder[any]{},
			false,
		},
		{
			"no cycles",
			map[Vertex[any]][]Vertex[any]{
				{"1", 1}: {Vertex[any]{"2", 2}},
				{"2", 2}: {Vertex[any]{"3", 3}, Vertex[any]{"4", 4}},
				{"3", 3}: {Vertex[any]{"4", 4}},
				{"4", 4}: {Vertex[any]{"5", 5}},
				{"5", 5}: {Vertex[any]{"6", 6}},
				{"6", 6}: {},
			},
			&TopoOrder[any]{Vertex[any]{"1", 1}, Vertex[any]{"2", 2}, Vertex[any]{"3", 3}, Vertex[any]{"4", 4}, Vertex[any]{"5", 5}, Vertex[any]{"6", 6}},
			false,
		},
		{
			"with cycles",
			map[Vertex[any]][]Vertex[any]{
				{"1", 1}: {Vertex[any]{"2", 2}},
				{"2", 2}: {Vertex[any]{"3", 3}, Vertex[any]{"4", 4}},
				{"3", 3}: {Vertex[any]{"4", 4}},
				{"4", 4}: {Vertex[any]{"5", 5}, Vertex[any]{"1", 1}},
				{"5", 5}: {Vertex[any]{"6", 6}},
				{"6", 6}: {},
			},
			nil,
			true,
		},
		{
			"with cycles and 2 edges",
			map[Vertex[any]][]Vertex[any]{
				{"1", 1}: {Vertex[any]{"2", 2}},
				{"2", 2}: {Vertex[any]{"1", 1}},
			},
			nil,
			true,
		},
		{
			"with cycles and self loop",
			map[Vertex[any]][]Vertex[any]{
				{"1", 1}: {Vertex[any]{"2", 2}, Vertex[any]{"1", 1}},
				{"2", 2}: {},
			},
			nil,
			true,
		},
		{
			"no cycles and job names",
			map[Vertex[any]][]Vertex[any]{
				{"clean", "clean"}:     {Vertex[any]{"build", "build"}},
				{"build", "build"}:     {Vertex[any]{"test", "test"}},
				{"test", "test"}:       {Vertex[any]{"deploy", "deploy"}},
				{"deploy", "deploy"}:   {Vertex[any]{"release", "release"}},
				{"release", "release"}: {},
			},
			&TopoOrder[any]{Vertex[any]{"clean", "clean"}, Vertex[any]{"build", "build"}, Vertex[any]{"test", "test"}, Vertex[any]{"deploy", "deploy"}, Vertex[any]{"release", "release"}},
			false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph[any]{}
			for v, edges := range tc.input {
				g.Add(v)
				for _, e := range edges {
					g.AddEdge(v, e, 1)
				}
			}
			actual, err := g.TopologicalSort()
			if tc.wantErr {
				require.Error(t, err)
				assert.ErrorContains(t, err, "there exists a cycle in the graph")
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})

	}
}
