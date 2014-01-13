package maneuver

import (
	"sync"
)

type EdgeSet struct {
	length uint64
	lock   sync.Mutex
	set    map[Edge]struct{}
}

func (n *EdgeSet) Add(edge Edge) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if n.Contains(edge) {
		return
	}
	n.set[edge] = struct{}{}
	n.length++
}

func (n *EdgeSet) Remove(edge Edge) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if !n.Contains(edge) {
		return
	}
	delete(n.set, edge)
	n.length--
}

func (n *EdgeSet) Contains(edge Edge) bool {
	_, ok := n.set[edge]
	return ok
}

func (n *EdgeSet) Empty() bool {
	return n.length == 0
}

func NewEdgeSet() *EdgeSet {
	return &EdgeSet{
		set: make(map[Edge]struct{}),
	}
}
