package dag

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTopologicalSort(t *testing.T) {
	cases := []struct {
		name     string
		input    map[any][]any
		expected *TopoOrder
		wantErr  bool
	}{
		{
			"empty graph",
			map[any][]any{},
			&TopoOrder{},
			false,
		},
		{
			"no cycles",
			map[any][]any{
				1: {2},
				2: {3, 4},
				3: {4},
				4: {5},
				5: {6},
				6: {},
			},
			&TopoOrder{1, 2, 3, 4, 5, 6},
			false,
		},
		{
			"with cycles",
			map[any][]any{
				1: {2},
				2: {3, 4},
				3: {4},
				4: {5, 1},
				5: {6},
				6: {},
			},
			nil,
			true,
		},
		{
			"with cycles and 2 edges",
			map[any][]any{
				1: {2},
				2: {1},
			},
			nil,
			true,
		},
		{
			"with cycles and self loop",
			map[any][]any{
				1: {1, 2},
				2: {},
			},
			nil,
			true,
		},
		{
			"no cycles and job names",
			map[any][]any{
				"clean":  {"build"},
				"build":  {"test"},
				"test":   {"deploy"},
				"deploy": {"release"},
			},
			&TopoOrder{"clean", "build", "test", "deploy", "release"},
			false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var g Graph
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
			assert.Equal(t, actual, tc.expected)
		})

	}
}
