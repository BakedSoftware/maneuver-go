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

func (f *FirstSearch) Path(grph *Graph, from, to Node, costAlgo uint8) []Node {
	nodesToCheck := make([]*firstSearchItem, 0)
	visited := make(map[*firstSearchItem]bool, 0)
	nodeToItem := make(map[Node]*firstSearchItem)
	fromSearchItem := firstSearchItem{from, nil}
	visited[&fromSearchItem] = true
	nodeToItem[from] = &fromSearchItem

	for _, edge := range from.OutgoingEdges() {
		item := firstSearchItem{edge.To, nodeToItem[edge.From]}
		nodesToCheck = append(nodesToCheck, &item)
	}

	for len(nodesToCheck) > 0 {
		var item *firstSearchItem
		if f.Method == BFS {
			item, nodesToCheck = nodesToCheck[len(nodesToCheck)-1], nodesToCheck[:len(nodesToCheck)-1]
		} else {
			item, nodesToCheck = nodesToCheck[0], nodesToCheck[1:len(nodesToCheck)]
		}
		if item.node == to {
			return f.reconstructPath(item)
		}
		for _, edge := range item.node.OutgoingEdges() {
			if _, ok := visited[nodeToItem[edge.To]]; !ok {
				test := firstSearchItem{edge.To, item}
				nodeToItem[edge.To] = &test
				visited[&test] = true
				nodesToCheck = append([]*firstSearchItem{&test}, nodesToCheck...)
			}
		}
	}
	return []Node{}
}
