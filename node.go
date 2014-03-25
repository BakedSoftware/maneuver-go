package maneuver

import (
	"strconv"
)

type Node interface {
	GetId() uint64
	Edges() *EdgeSet
	SetEdges(e *EdgeSet)
	// Get the edges that originate from the node and are not re-entrant on the
	// node. Must be free of single node cycles (Edge.From != Edge.To)
	OutgoingEdges() []Edge
	Clone(c *Cloner) Node
	// Used when create DOT output
	String() string
}

type GraphNode struct {
	Id    uint64
	edges *EdgeSet
}

func (g *GraphNode) OutgoingEdges() []Edge {
	edges := make([]Edge, 0)
	for e := range g.edges.AllEdges() {
		if e.ToNode().GetId() != g.Id {
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

func (g *GraphNode) String() string {
	return strconv.Itoa(int(g.Id))
}

func (g *GraphNode) Clone(c *Cloner) Node {
	if c.Nodes.Contains(g) {
		return c.Nodes.Get(g.GetId())
	}
	clone := &GraphNode{
		Id:    g.Id,
		edges: g.edges,
	}
	c.Nodes.Add(clone)
	return clone
}
