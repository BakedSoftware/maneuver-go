package maneuver

type firstSearchItem struct {
	node   Node
	parent *firstSearchItem
}

type FirstSearch struct {
	Method uint8
}

func (f *FirstSearch) reconstructPath(item *firstSearchItem) []Node {
	path := make([]Node, 0)
	for item.parent != nil {
		path = append([]Node{item.node}, path...)
		item = item.parent
	}
	path = append([]Node{item.node}, path...)
	return path
}

func (f *FirstSearch) Path(graph *Graph, from, to Node, costAlgo uint8) []Node {
	nodesToCheck := make([]*firstSearchItem, 0)
	visited := make(map[*firstSearchItem]bool, 0)
	nodeToItem := make(map[uint64]*firstSearchItem)
	fromSearchItem := firstSearchItem{from, nil}
	visited[&fromSearchItem] = true
	nodeToItem[from.GetId()] = &fromSearchItem

	for _, edge := range graph.OutgoingEdgesForNode(from) {
		item := firstSearchItem{edge.ToNode(), nodeToItem[edge.FromNode().GetId()]}
		nodesToCheck = append(nodesToCheck, &item)
	}

	for len(nodesToCheck) > 0 {
		var item *firstSearchItem
		if f.Method == BFS {
			item, nodesToCheck = nodesToCheck[len(nodesToCheck)-1], nodesToCheck[:len(nodesToCheck)-1]
		} else {
			item, nodesToCheck = nodesToCheck[0], nodesToCheck[1:len(nodesToCheck)]
		}
		if item.node.GetId() == to.GetId() {
			return f.reconstructPath(item)
		}
		for _, edge := range graph.OutgoingEdgesForNode(item.node) {
			if _, ok := visited[nodeToItem[edge.ToNode().GetId()]]; !ok {
				test := firstSearchItem{edge.ToNode(), item}
				nodeToItem[edge.ToNode().GetId()] = &test
				visited[&test] = true
				nodesToCheck = append([]*firstSearchItem{&test}, nodesToCheck...)
			}
		}
	}
	return []Node{}
}
