package maneuver

import (
	"fmt"
	"log"
)

type Graph struct {
	Nodes *NodeSet
	Edges *EdgeSet
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: NewNodeSet(),
		Edges: NewEdgeSet(),
	}
}

func NewGraphWithEdges(edges *EdgeSet) *Graph {
	graph := Graph{Edges: edges, Nodes: NewNodeSet()}
	for e := range edges.set {
		if !graph.ContainsNode(e.FromNode()) {
			graph.AddNode(e.FromNode())
		}
		if !graph.ContainsNode(e.ToNode()) {
			graph.AddNode(e.ToNode())
		}
	}
	return &graph
}

func (g *Graph) ContainsNode(node Node) bool {
	return g.Nodes.Contains(node)
}

func (g *Graph) AddNode(node Node) {
	g.Nodes.Add(node)
}

func (g *Graph) InsertEdges(edges ...Edge) {
	for _, e := range edges {
		if !g.ContainsMatchingEdge(e) {
			g.Edges.Add(e)
			if !g.ContainsNode(e.FromNode()) {
				g.AddNode(e.FromNode())
			}
			if !g.ContainsNode(e.ToNode()) {
				g.AddNode(e.ToNode())
			}
		}
	}
}

func (g *Graph) ContainsEdge(edge Edge) bool {
	return g.Edges.Contains(edge)
}

func (g *Graph) ContainsMatchingEdge(edge Edge) bool {
	if yes := g.ContainsEdge(edge); yes {
		return true
	}

	return g.EdgeBetween(edge.FromNode(), edge.ToNode()) != nil
}

func (g *Graph) Path(from, to Node, searchKey, costKey uint8) []Node {
	if from.GetId() == to.GetId() {
		return []Node{from}
	}
	search := GetPathAlgorithm(searchKey)
	if search == nil {
		log.Panicln("Unknown Path Algorithm: ", searchKey)
	}
	return search.Path(g, g.GetNodeWithId(from.GetId()), g.GetNodeWithId(to.GetId()), costKey)
}

func (g *Graph) EdgeBetween(from, to Node) Edge {
	for _, e := range g.OutgoingEdgesForNode(from) {
		if e.ToNode().GetId() == to.GetId() {
			return e
		}
	}
	return nil
}

func (g *Graph) OutgoingEdgesForNode(node Node) []Edge {
	out := node.OutgoingEdges()
	edges := make([]Edge, 0, len(out))
	for _, e := range out {
		if g.Edges.Contains(e) {
			edges = append(edges, e)
		}
	}
	return edges
}

func (g *Graph) GetNodeWithId(id uint64) Node {
	return g.Nodes.set[id]
}

func (g *Graph) Clone() *Graph {
	edges := NewEdgeSet()
	for e, _ := range g.Edges.set {
		edges.Add(e.Clone())
	}
	return NewGraphWithEdges(edges)
}

func (g *Graph) DOT() string {
	dot := "digraph G {\n"
	for e := range g.Edges.AllEdges() {
		dot += fmt.Sprintf("\t%s;\n", e.String())
	}
	dot += "}"
	return dot
}
