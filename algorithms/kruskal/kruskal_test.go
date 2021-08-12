package kruskal

import (
	. "github.com/judimator/algorithms/datastructures/graph"
	"reflect"
	"testing"
)

func TestGetMinimumSpanningTree(t *testing.T) {
	if !reflect.DeepEqual(getExpectedTree(), GetMinimumSpanningTree(getGraph())) {
		t.Error("Invalid graph received.")
	}
}

func getGraph() *Graph {
	graph := New()
	graph.AddEdge(0, 1, 3)
	graph.AddEdge(1, 2, 5)
	graph.AddEdge(2, 3, 4)
	graph.AddEdge(3, 4, 1)
	graph.AddEdge(4, 5, 2)
	graph.AddEdge(3, 0, 1)

	return graph
}

func getExpectedTree() *Graph {
	graph := New()
	graph.AddEdge(3, 4, 1)
	graph.AddEdge(3, 0, 1)
	graph.AddEdge(4, 5, 2)
	graph.AddEdge(0, 1, 3)
	graph.AddEdge(2, 3, 4)

	return graph
}
