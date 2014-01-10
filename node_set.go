package maneuver

import (
	"sync"
)

type NodeSet struct {
	length uint64
	lock   sync.Mutex
	set    map[Node]struct{}
}

func (n *NodeSet) Add(node Node) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if n.Contains(node) {
		return
	}
	n.set[node] = struct{}{}
	n.length++
}

func (n *NodeSet) Remove(node Node) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if !n.Contains(node) {
		return
	}
	delete(n.set, node)
	n.length--
}

func (n *NodeSet) Contains(node Node) bool {
	_, ok := n.set[node]
	return ok
}

func (n *NodeSet) Empty() bool {
	return n.length == 0
}

func NewNodeSet() *NodeSet {
	return &NodeSet{
		set: make(map[Node]struct{}),
	}
}
