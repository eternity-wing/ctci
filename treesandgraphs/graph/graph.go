package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	Vertices map[string]Vertex
}

func CreateNewGraph(vertices []string) Graph {
	graph := Graph{Vertices: make(map[string]Vertex)}
	for _, vertex := range vertices {
		graph.Vertices[vertex] = NewVertex(vertex)
	}
	return graph
}

func (g *Graph) AddEdge(v1 string, v2 string) {
	v1Node, found1 := g.Vertices[v1]
	if !found1 {
		v1Node = NewVertex(v1)
		g.Vertices[v1] = v1Node
	}

	v2Node, found2 := g.Vertices[v2]
	if !found2 {
		v2Node = NewVertex(v2)
		g.Vertices[v2] = v2Node
	}

	v1Node.AddNeighbour(v2Node)
}

func (g *Graph) PrintTopologicalSort() {
	pathStack := list.New()
	visitedVertices := make(map[string]bool)
	for _, vertex := range g.Vertices {
		g.TopologicalSort(vertex, visitedVertices, pathStack)
	}
	if pathStack.Len() != len(g.Vertices) {
		fmt.Printf("There is a cycle in graph")
	}
	path := ""
	for vertex := pathStack.Front(); vertex != nil; vertex = vertex.Next() {
		path += vertex.Value.(Vertex).Data
	}
	fmt.Printf("%s", path)
}
func (g *Graph) TopologicalSort(vertex Vertex, visitedVertices map[string]bool, pathStack *list.List) {
	if _, isVisited := visitedVertices[vertex.Data]; isVisited {
		return
	}
	visitedVertices[vertex.Data] = true

	for _, neighbourVertex := range vertex.Neighbours {
		g.TopologicalSort(neighbourVertex, visitedVertices, pathStack)
	}
	pathStack.PushFront(vertex)
}
