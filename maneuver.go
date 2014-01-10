// Package maneuver provides a flexible library of common path planning
// algorithms. The design focus is on making it easy to quickly add and test
// new weighting and planning algorithms.
package maneuver

// Interface to implement to add a new cost algorithm
type CostAlgorithm interface {
	Cost(from, to Node) float64
}

// Interface to implement to add a new path finding algorithm
type PathAlgorithm interface {
	Path(graph *Graph, from, to Node, costAlgo uint8) []Node
}

var pathAlgorithms map[uint8]PathAlgorithm = make(map[uint8]PathAlgorithm)
var costAlgorithms map[uint8]CostAlgorithm = make(map[uint8]CostAlgorithm)

// Add a path aglorithm to Maneuver for graphs to use
func RegisterPathAlgorithm(a PathAlgorithm, name uint8) {
	pathAlgorithms[name] = a
}

// Retrieve a path algorithm
func GetPathAlgorithm(name uint8) PathAlgorithm {
	return pathAlgorithms[name]
}

// Add a cost algorithm to Maneuver for planning algorithms that require
// weighting. i.e A Star
func RegisterCostAlgorithm(a CostAlgorithm, name uint8) {
	costAlgorithms[name] = a
}

// Retrieve a cost algorithm
func GetCostAlgorithm(name uint8) CostAlgorithm {
	return costAlgorithms[name]
}

const (
	// Breadth First Search
	BFS = iota
	// Depth First Search
	DFS
	// A Star Search
	AStar
	// Assign new path algorithms keys equal to or above
	// Offset value may change so use UserSearchKeyOffset as iota
	UserSearchKeyOffset
)

const (
	// Don't use any cost algorithm - BFS and DFS
	NONE = iota
	// Assign new cost algorithms keys equal to or above
	// Offset value may change so use UserCostKeyOffset as iota
	UserCostKeyOffset
)

func init() {
	RegisterPathAlgorithm(&FirstSearch{Method: BFS}, BFS)
	RegisterPathAlgorithm(&FirstSearch{Method: DFS}, DFS)
	RegisterPathAlgorithm(&AStarSearch{}, AStar)
}
