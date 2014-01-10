package maneuver

type Edge struct {
	From Node
	To   Node
}

func NewEdge(from, to Node) *Edge {
	e := Edge{from, to}
	if from.Edges() == nil {
		from.SetEdges(NewEdgeSet())
	}
	if to.Edges() == nil {
		to.SetEdges(NewEdgeSet())
	}
	from.Edges().Add(&e)
	to.Edges().Add(&e)

	return &e
}
