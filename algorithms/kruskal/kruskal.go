package kruskal

import (
	"github.com/judimator/algorithms/algorithms/unionfind"
	"github.com/judimator/algorithms/datastructures/graph"
	"sort"
)

type EdgeHeap []graph.Edge

func (e EdgeHeap) Len() int {
	return len(e)
}

func (e EdgeHeap) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e EdgeHeap) Less(i, j int) bool {
	return e[i].Weight() < e[j].Weight()
}

func GetMinimumSpanningTree(g *graph.Graph) *graph.Graph {
	sort.Sort(EdgeHeap(g.Edges()))

	set := unionfind.New(len(g.Edges()))
	newGraph := graph.New()

	for _, edge := range g.Edges() {
		from, to := edge.Vertices()

		uRoot := set.Find(from.Id())
		vRoot := set.Find(to.Id())
		if uRoot != vRoot {
			set.Union(uRoot, vRoot)
			newGraph.AddEdge(from.Id(), to.Id(), edge.Weight())
		}
	}

	return newGraph
}
