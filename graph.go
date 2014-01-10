package maneuver

import (
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
	for _, e := range edges.set {
		if !graph.ContainsNode(e.From) {
			graph.AddNode(e.From)
		}
		if !graph.ContainsNode(e.To) {
			graph.AddNode(e.To)
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

func (g *Graph) Path(from, to Node, searchKey, costKey uint8) []Node {
	if from == to {
		return []Node{from}
	}
	search := GetPathAlgorithm(searchKey)
	if search == nil {
		log.Panicln("Unknown Path Algorithm: ", searchKey)
	}
	return search.Path(g, from, to, costKey)
}

func (g *Graph) EdgeBetween(from, to Node) *Edge {
	for _, e := range from.OutgoingEdges() {
		if e.To == to && g.Edges.Contains(e) {
			return e
		}
	}
	return nil
}

// func (g *Graph) Clone() *Graph {
// 	edges := NewEdgeSet()
// 	for e, _ := range g.Edges.set {
// 		edges.Add(&Edge{
// 			From:   e.From.Clone(),
// 			To:     e.To.Clone(),
// 			Weight: e.Weight,
// 		})
// 	}
// 	return NewGraphWithEdges(edges)
// }
