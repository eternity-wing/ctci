package graph

type Node struct {
	Neighbors map[string]Node
	Data      string
}

func NewNode(data string) Node {
	return Node{Data: data, Neighbors: make(map[string]Node)}
}

func (n *Node) IsNeighbor(v Node) bool {
	_, found := n.Neighbors[v.Data]
	return found
}

func (n *Node) AddNeighbor(v Node)  {
	if !n.IsNeighbor(v){
		n.Neighbors[v.Data] = v
	}
}