package maneuver

type Node interface {
	GetId() uint64
	Edges() *EdgeSet
	SetEdges(e *EdgeSet)
	// Get the edges that originate from the node and are not re-entrant on the
	// node. Must be free of single node cycles (Edge.From != Edge.To)
	OutgoingEdges() []*Edge
}

type GraphNode struct {
	Id    uint64
	edges *EdgeSet
}

func (g *GraphNode) OutgoingEdges() []*Edge {
	edges := make([]*Edge, 0)
	for _, e := range g.edges.set {
		if e.To.GetId() != g.Id {
			edges = append(edges, e)
		}
	}
	return edges
}

func (g *GraphNode) GetId() uint64 {
	return g.Id
}

func (g *GraphNode) Edges() *EdgeSet {
	return g.edges
}

func (g *GraphNode) SetEdges(e *EdgeSet) {
	g.edges = e
}
