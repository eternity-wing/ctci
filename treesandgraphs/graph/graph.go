package graph



type Graph struct {
	Nodes map[string]Node
}

func CreateNewGraph() Graph{
	return Graph{Nodes: make(map[string]Node)}
}

func (g *Graph) AddEdge(v1 string, v2 string) {
	v1Node, found1 := g.Nodes[v1]
	if !found1{
		v1Node = NewNode(v1)
		g.Nodes[v1] = v1Node
	}

	v2Node, found2 := g.Nodes[v2]
	if !found2{
		v2Node = NewNode(v2)
		g.Nodes[v2] = v2Node
	}

	v1Node.AddNeighbor(v2Node)
}
