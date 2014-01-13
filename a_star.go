package maneuver

import (
	"math"
)

type AStarSearch struct {
}

func (a *AStarSearch) Path(grph *Graph, from, to Node, costAlgo uint8) []Node {
	if costAlgo == NONE {
		panic("Edge cost algorithm must be set")
	}

	open := NewNodeSet()
	open.Add(from)
	estimate := GetCostAlgorithm(costAlgo)
	cameFrom := make(map[Node]Node, 0)
	gScore := map[Node]float64{from: 0.0}
	fScore := map[Node]float64{from: estimate.Cost(from, to)}

	for !open.Empty() {
		current := a.keyWithMinValue(open, fScore)
		if current == to {
			return a.reconstructPath(cameFrom, to)
		}
		open.Remove(current)
		for _, e := range current.OutgoingEdges() {
			n := e.ToNode()
			tGScore := gScore[current] + estimate.Cost(e.FromNode(), n)
			if ok := open.Contains(n); !ok || tGScore <= gScore[n] {
				cameFrom[n] = current
				gScore[n] = tGScore
				fScore[n] = tGScore + estimate.Cost(n, to)
				if !ok {
					open.Add(n)
				}
			}
		}
	}

	return make([]Node, 0)
}

func (a *AStarSearch) keyWithMinValue(set *NodeSet, hash map[Node]float64) Node {
	minValue := math.MaxFloat64
	var minKey Node = nil
	for _, k := range set.set {
		m := hash[k]
		if m < minValue {
			minKey = k
			minValue = m
		}
	}
	return minKey
}

func (a *AStarSearch) reconstructPath(cameFrom map[Node]Node, current Node) []Node {
	path := make([]Node, 1)
	path[0] = current
	node, ok := cameFrom[current]
	for ok {
		path = append([]Node{node}, path...)
		current = node
		node, ok = cameFrom[current]
	}
	return path
}
