package maneuver

import (
	"fmt"
	"sync"
)

// Store a set of edges. Edges are determined as unique based on the pointer
// values of Edge.From and Edge.To
type EdgeSet struct {
	length uint64
	lock   sync.Mutex
	set    map[string]*Edge
}

func (n *EdgeSet) Add(edge *Edge) {
	n.lock.Lock()
	defer n.lock.Unlock()
	key := keyForEdge(edge)
	if n.containsKey(key) {
		return
	}
	n.set[key] = edge
	n.length++
}

func (n *EdgeSet) Remove(edge *Edge) {
	n.lock.Lock()
	defer n.lock.Unlock()
	key := keyForEdge(edge)
	if !n.containsKey(key) {
		return
	}
	delete(n.set, key)
	n.length--
}

func (n *EdgeSet) containsKey(key string) bool {
	_, ok := n.set[key]
	return ok
}

func (n *EdgeSet) Contains(edge *Edge) bool {
	return n.containsKey(keyForEdge(edge))
}

func keyForEdge(edge *Edge) string {
	return fmt.Sprintf("%p-%p", edge.From, edge.To)
}

func (n *EdgeSet) Empty() bool {
	return n.length == 0
}

func NewEdgeSet() *EdgeSet {
	return &EdgeSet{
		set: make(map[string]*Edge),
	}
}
