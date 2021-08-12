package graph

type Vertex struct {
	id               int
	adjacentVertices []*Vertex
}

type Edge struct {
	from   *Vertex
	to     *Vertex
	weight int
}

type Graph struct {
	edges    []Edge
	vertices map[int]*Vertex
}

func (g *Graph) AddEdge(from, to, weight int) {
	fromVertex := g.addVertex(from)
	toVertex := g.addVertex(to)

	fromVertex.adjacentVertices = append(fromVertex.adjacentVertices, toVertex)
	toVertex.adjacentVertices = append(toVertex.adjacentVertices, fromVertex)

	e := Edge{
		from:   fromVertex,
		to:     toVertex,
		weight: weight,
	}

	g.edges = append(g.edges, e)
}

func (g *Graph) Edges() []Edge {
	return g.edges
}

func (g *Graph) addVertex(v int) *Vertex {
	if _, ok := g.vertices[v]; !ok {
		g.vertices[v] = &Vertex{
			id: v,
		}
	}

	return g.vertices[v]
}

func (g *Graph) Vertices() map[int]*Vertex {
	return g.vertices
}

func (v *Vertex) AdjacentVertices() []*Vertex {
	return v.adjacentVertices
}

func (v *Vertex) Id() int {
	return v.id
}

func (e *Edge) Weight() int {
	return e.weight
}

func (e *Edge) Vertices() (*Vertex, *Vertex) {
	return e.from, e.to
}

func New() *Graph {
	return &Graph{
		vertices: make(map[int]*Vertex),
	}
}
