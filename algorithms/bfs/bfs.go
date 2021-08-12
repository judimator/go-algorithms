package bfs

import (
	"fmt"
	"github.com/judimator/algorithms/datastructures/graph"
)
import "github.com/judimator/algorithms/datastructures/queue"

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

	q := queue.New()
	q.Push(v)

	for {
		if q.Len() == 0 {
			break
		}

		v = q.Pop().(*graph.Vertex)

		if _, ok := visitedNodes[v.Id()]; !ok {
			ps = append(ps, v.Id())
		}

		visitedNodes[v.Id()] = true

		for _, v := range v.AdjacentVertices() {
			if _, ok := visitedNodes[v.Id()]; !ok {
				q.Push(v)
			}
		}
	}

	return ps, nil
}
