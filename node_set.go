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

func (n *NodeSet) Get(id uint64) Node {
	return n.set[id]
}

func (n *NodeSet) AllNodes() []Node {
	nodes := make([]Node, len(n.set))
	i := 0
	for _, v := range n.set {
		nodes[i], i = v, i+1
	}
	return nodes
}

func (n *NodeSet) Size() uint64 {
	return n.length
}

func NewNodeSet() *NodeSet {
	return &NodeSet{
		set: make(map[uint64]Node),
	}
}
