package maneuver

import (
	"sync"
)

type NodeSet struct {
	length uint64
	lock   sync.Mutex
	set    map[uint64]Node
}

func (n *NodeSet) Add(node Node) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if n.Contains(node) {
		return
	}
	n.set[node.GetId()] = node
	n.length++
}

func (n *NodeSet) Remove(node Node) {
	n.lock.Lock()
	defer n.lock.Unlock()
	if !n.Contains(node) {
		return
	}
	delete(n.set, node.GetId())
	n.length--
}

func (n *NodeSet) Contains(node Node) bool {
	_, ok := n.set[node.GetId()]
	return ok
}

func (n *NodeSet) Empty() bool {
	return n.length == 0
}

func NewNodeSet() *NodeSet {
	return &NodeSet{
		set: make(map[uint64]Node),
	}
}
