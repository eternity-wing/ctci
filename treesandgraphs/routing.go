package treesandgraphs

import (
	"container/list"
	"github.com/eternity-wings/dev_talks/treesandgraphs/graph"
)

func DoesRouteExist(g graph.Graph, source string, destination string) bool {
	sourceNode := g.Vertices[source]
	destinationNode := g.Vertices[destination]

	var visitedNodes = make(map[string]bool)
	path := list.New()

	path.PushBack(sourceNode)
	visitedNodes[sourceNode.Data] = true

	for frontElement := path.Front(); path.Len() > 0; frontElement = path.Front() {
		frontNode := frontElement.Value.(graph.Vertex)
		path.Remove(frontElement)
		if frontNode.Data == destinationNode.Data {
			return true
		}
		visitUnvisitedNeighborsOfTheNode(frontNode, path, visitedNodes)
	}
	return false
}


func visitUnvisitedNeighborsOfTheNode(node graph.Vertex, visitQueue *list.List, visitedNodes map[string]bool)  {
	for _, nextNode := range node.Neighbours {
		if !visitedNodes[nextNode.Data] {
			visitQueue.PushBack(nextNode)
			visitedNodes[nextNode.Data] = true
		}
	}
}
