package maneuver

import (
	"fmt"
	"math"
	"testing"
)

type TransportNode struct {
	GraphNode
	Lat, Lon float64
}

type ManhattanDistance struct {
}

func (m *ManhattanDistance) Cost(from, to Node) float64 {
	f := from.(*TransportNode)
	t := to.(*TransportNode)
	lat_d := f.Lat - t.Lat
	lon_d := f.Lon - t.Lon
	return lat_d*lat_d + lon_d*lon_d
}

func ExampleRegisterCostAlgorithm() {

	// Having defined a cost algorithm (struct) as follows

	// type ManhattanDistance struct {
	// }

	// func (m *ManhattanDistance) Cost(from, to Node) float64 {
	// 	f := from.(*TransportNode)
	// 	t := to.(*TransportNode)
	// 	lat_d := f.Lat - t.Lat
	// 	lon_d := f.Lon - t.Lon
	// 	return lat_d*lat_d + lon_d*lon_d
	// }

	// const (
	// 	 Manhattan = UserCostKeyOffset
	// )

	// It can registered as available using:
	RegisterCostAlgorithm(&ManhattanDistance{}, name)

	// Then you can call graph.Path(from, to, AStar, Manhattan) and a path
	// between from and to will be found using A Star and Manhattan distance
}

type EuclidianDistance struct {
}

func (e *EuclidianDistance) Cost(from, to *TransportNode) float64 {
	return math.Sqrt((&ManhattanDistance{}).Cost(from, to))
}

const (
	Manhattan = UserCostKeyOffset
	Euclidian
)

func ExampleNewGraphWithEdges() {
	// GraphNode is a builtin type that has base implementations of the Node
	// interface. TransportNode is a user defined type with additional fields
	n1 := &TransportNode{GraphNode: GraphNode{Id: 1}, Lat: 5, Lon: 5}
	n2 := &TransportNode{GraphNode: GraphNode{Id: 2}, Lat: 10, Lon: 10}

	edges := NewEdgeSet()
	edges.Add(NewEdge(n1, n2))

	/* graph := */ NewGraphWithEdges(edges)

}

var n1 = &TransportNode{GraphNode: GraphNode{Id: 1}, Lat: 5, Lon: 5}
var n2 = &TransportNode{GraphNode: GraphNode{Id: 2}, Lat: 10, Lon: 10}
var n3 = &TransportNode{GraphNode: GraphNode{Id: 3}, Lat: 15, Lon: 15}
var n4 = &TransportNode{GraphNode: GraphNode{Id: 4}, Lat: 20, Lon: 20}
var n5 = &TransportNode{GraphNode: GraphNode{Id: 5}, Lat: 25, Lon: 25}
var n6 = &TransportNode{GraphNode: GraphNode{Id: 6}, Lat: 900, Lon: 900}
var n7 = &TransportNode{GraphNode: GraphNode{Id: 7}, Lat: 40, Lon: 40}

func buildGraph() *Graph {
	edges := NewEdgeSet()
	edges.Add(NewEdge(n1, n2))
	edges.Add(NewEdge(n2, n3))
	edges.Add(NewEdge(n1, n4))
	edges.Add(NewEdge(n4, n5))
	edges.Add(NewEdge(n6, n7))
	edges.Add(NewEdge(n1, n6))
	edges.Add(NewEdge(n3, n7))

	return NewGraphWithEdges(edges)
}

func Example() {

	// Create some nodes for the graph
	// TransportNode is a user defined type adding Lat and Lon fields
	// GraphNode is a builtin type that implements the Node interface
	n1 := &TransportNode{GraphNode: GraphNode{Id: 1}, Lat: 5, Lon: 5}
	n2 := &TransportNode{GraphNode: GraphNode{Id: 2}, Lat: 10, Lon: 10}
	n3 := &TransportNode{GraphNode: GraphNode{Id: 3}, Lat: 15, Lon: 15}
	n4 := &TransportNode{GraphNode: GraphNode{Id: 4}, Lat: 20, Lon: 20}
	n5 := &TransportNode{GraphNode: GraphNode{Id: 5}, Lat: 25, Lon: 25}
	n6 := &TransportNode{GraphNode: GraphNode{Id: 6}, Lat: 900, Lon: 900}
	n7 := &TransportNode{GraphNode: GraphNode{Id: 7}, Lat: 40, Lon: 40}

	// Create the edges between the nodes. Edges are directed.
	// Always use NewEdgeSet and NewEdge.
	edges := NewEdgeSet()
	edges.Add(NewEdge(n1, n2))
	edges.Add(NewEdge(n2, n3))
	edges.Add(NewEdge(n1, n4))
	edges.Add(NewEdge(n4, n5))
	edges.Add(NewEdge(n6, n7))
	edges.Add(NewEdge(n1, n6))
	edges.Add(NewEdge(n3, n7))

	// Create a new graph
	graph := NewGraphWithEdges(edges)

	// Get a path
	path := graph.Path(n1, n7, BFS, NONE)

	for idx, node := range path {
		fmt.Printf("%d", node.GetId())
		if idx+1 < len(path) {
			fmt.Print("->")
		}
	}
	// Output:
	// 1->6->7

}

func TestBreadthFirstSearch(t *testing.T) {
	graph := buildGraph()
	path := graph.Path(n1, n7, BFS, NONE)

	//Expected: n1 -> n6 -> n7
	if len(path) != 3 {
		t.Log("Incorrect path length. Expected 3 got ", len(path))
		for _, n := range path {
			t.Log(n.(*TransportNode).GetId())
		}
		t.FailNow()
	}
	var node *TransportNode
	node = path[0].(*TransportNode)
	if node.GetId() != n1.GetId() || node != n1 {
		t.Fatal("First Node must be From node, got", node.GetId())
	}
	node = path[1].(*TransportNode)
	if node.GetId() != n6.GetId() || node != n6 {
		t.Fatal("Path should have gone through N6, got", node.GetId())
	}
	node = path[2].(*TransportNode)
	if node.GetId() != n7.GetId() || node != n7 {
		t.Fatal("Last Node must be To node, got", node.GetId())
	}
}

func TestDepthFirstSearch(t *testing.T) {
	graph := buildGraph()
	path := graph.Path(n1, n7, DFS, NONE)

	//Expected: n1 -> n2 -> n3 -> n7
	if len(path) != 4 {
		t.Log("Incorrect path length. Expected 4 got ", len(path))
		for _, n := range path {
			t.Log(n.(*TransportNode).GetId())
		}
		t.FailNow()
	}
	var node *TransportNode
	node = path[0].(*TransportNode)
	if node.GetId() != n1.GetId() || node != n1 {
		t.Fatal("First Node must be From node, got", node.GetId())
	}
	node = path[1].(*TransportNode)
	if node.GetId() != n2.GetId() || node != n2 {
		t.Fatal("Path should have gone through N2, got", node.GetId())
	}
	node = path[2].(*TransportNode)
	if node.GetId() != n3.GetId() || node != n3 {
		t.Fatal("Path should have gone through N3, got", node.GetId())
	}
	node = path[3].(*TransportNode)
	if node.GetId() != n7.GetId() || node != n7 {
		t.Fatal("Last Node must be To node, got", node.GetId())
	}
}

func TestAStarSearch(t *testing.T) {
	graph := buildGraph()
	RegisterCostAlgorithm(&ManhattanDistance{}, Manhattan)
	path := graph.Path(n1, n7, AStar, Manhattan)

	//Expected: n1 -> n2 -> n3 -> n7
	if len(path) != 4 {
		t.Log("Incorrect path length. Expected 4 got ", len(path))
		for _, n := range path {
			t.Log(n.(*TransportNode).GetId())
		}
		t.FailNow()
	}
	var node *TransportNode
	node = path[0].(*TransportNode)
	if node.GetId() != n1.GetId() || node != n1 {
		t.Fatal("First Node must be From node, got", node.GetId())
	}
	node = path[1].(*TransportNode)
	if node.GetId() != n2.GetId() || node != n2 {
		t.Fatal("Path should have gone through N2, got", node.GetId())
	}
	node = path[2].(*TransportNode)
	if node.GetId() != n3.GetId() || node != n3 {
		t.Fatal("Path should have gone through N3, got", node.GetId())
	}
	node = path[3].(*TransportNode)
	if node.GetId() != n7.GetId() || node != n7 {
		t.Fatal("Last Node must be To node, got", node.GetId())
	}
}
