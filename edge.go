package maneuver

import (
	"fmt"
)

type GraphEdge struct {
	From Node
	To   Node
}

type Edge interface {
	FromNode() Node
	ToNode() Node
	Clone() Edge
	String() string
}

func (g *GraphEdge) Clone() Edge {
	return NewGraphEdge(g.From.Clone(), g.To.Clone())
}

func (g *GraphEdge) FromNode() Node {
	return g.From
}

func (g *GraphEdge) ToNode() Node {
	return g.To
}

func (g *GraphEdge) String() string {
	return fmt.Sprintf("%s -> %s", g.From.String(), g.To.String())
}

func NewGraphEdge(from, to Node) *GraphEdge {
	e := GraphEdge{from, to}
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
