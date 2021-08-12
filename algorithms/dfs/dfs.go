package dfs

import (
	"fmt"
	"github.com/judimator/algorithms/datastructures/graph"
)
import "github.com/judimator/algorithms/datastructures/stack"

type errorNodeNotFound struct {
	id int
}

func (e errorNodeNotFound) Error() string {
	return fmt.Sprintf("Error with %d not found", e.id)
}

func Pass(g *graph.Graph, id int) ([]int, error) {
	vs := g.Vertices()
	v, ok := vs[id]

	if !ok {
		return nil, errorNodeNotFound{id}
	}
	// The path algorithm passes through
	var ps []int

	visitedNodes := make(map[int]bool)

	s := stack.New()
	s.Push(v)

	for {
		if s.Len() == 0 {
			break
		}

		v = s.Pop().(*graph.Vertex)

		if _, ok := visitedNodes[v.Id()]; !ok {
			ps = append(ps, v.Id())
		}

		visitedNodes[v.Id()] = true

		for _, v := range v.AdjacentVertices() {
			if _, ok := visitedNodes[v.Id()]; !ok {
				s.Push(v)
			}
		}
	}

	return ps, nil
}
