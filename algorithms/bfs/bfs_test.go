package bfs

import (
	. "github.com/judimator/algorithms/datastructures/graph"
	"reflect"
	"testing"
)

func TestPass(t *testing.T) {
	path, _ := Pass(getGraph(), 0)
	if !reflect.DeepEqual(getExpectedPass(), path) {
		t.Error("Invalid route received.")
	}
}

func getExpectedPass() []int {
	return []int{0, 1, 3, 2, 7, 4, 6, 8, 5}
}

func getGraph() *Graph {
	graph := New()
	graph.AddEdge(0, 1, 0)
	graph.AddEdge(1, 2, 0)
	graph.AddEdge(2, 3, 0)
	graph.AddEdge(3, 4, 0)
	graph.AddEdge(4, 5, 0)
	graph.AddEdge(3, 0, 0)
	graph.AddEdge(1, 7, 0)
	graph.AddEdge(2, 6, 0)
	graph.AddEdge(2, 8, 0)
	graph.AddEdge(6, 8, 0)
	graph.AddEdge(4, 6, 0)

	return graph
}
