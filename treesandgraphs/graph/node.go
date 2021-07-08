package graph

type Vertex struct {
	Neighbours map[string]Vertex
	Data       string
}

func NewVertex(data string) Vertex {
	return Vertex{Data: data, Neighbours: make(map[string]Vertex)}
}

func (n *Vertex) IsNeighbour(v Vertex) bool {
	_, found := n.Neighbours[v.Data]
	return found
}

func (n *Vertex) AddNeighbour(v Vertex)  {
	if !n.IsNeighbour(v){
		n.Neighbours[v.Data] = v
	}
}