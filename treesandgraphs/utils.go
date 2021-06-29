package treesandgraphs

import (
	"container/list"
	"fmt"
	"math"
)

func PrintInOrder(node *BSTNode) {
	if node == nil {
		return
	}
	PrintInOrder(node.LeftChild)
	fmt.Printf("%d -> ", node.Data)
	PrintInOrder(node.RightChild)
}

func GenerateMinimalTree(t *BinarySearchTree, array []int, start int, end int) {
	if start > end {
		return
	}
	rootIndex := start + ((end - start) / 2)
	t.Add(array[rootIndex])
	GenerateMinimalTree(t, array, start, rootIndex-1)
	GenerateMinimalTree(t, array, rootIndex+1, end)
}


func GetBSTNodeBalanceFactor(node BSTNode) int {
	return GetBSTNodeHeight(node.LeftChild) - GetBSTNodeHeight(node.RightChild)
}

func GetBSTNodeHeight(node *BSTNode) int {
	if node == nil{
		return -1
	}
	return int(math.Max(float64(GetBSTNodeHeight(node.LeftChild)), float64(GetBSTNodeHeight(node.RightChild))) + 1)
}


func GetBreathNodes(bst BinarySearchTree) map[int]*list.List {
	listOfDepths := make(map[int]*list.List)
	firstDepthNodes := list.New()
	firstDepthNodes.PushBack(bst.Root)
	StoreDepthNodesList(firstDepthNodes, 0, listOfDepths)
	return listOfDepths
}


func StoreDepthNodesList(sameLevelNodes *list.List, level int, listOfDepths map[int]*list.List)  {
	if sameLevelNodes == nil || sameLevelNodes.Len() == 0{
		return
	}
	listOfDepths[level] = sameLevelNodes
	StoreDepthNodesList(GetNodesListChildren(sameLevelNodes), level+1, listOfDepths)
}


func GetNodesListChildren(nodes *list.List) *list.List {
	childrenQueue := list.New()
	for frontElement := nodes.Front(); frontElement != nil; frontElement = frontElement.Next() {
		EnqueueNodeChildren(frontElement.Value.(*BSTNode), childrenQueue)
	}
	return childrenQueue
}


func EnqueueNodeChildren(node *BSTNode, queue *list.List)  {
	if node == nil{
		return
	}
	if node.LeftChild != nil {
		queue.PushBack(node.LeftChild)
	}

	if node.RightChild != nil {
		queue.PushBack(node.RightChild)
	}
}


func PrintBreathNodes(bst BinarySearchTree)  {
	result := GetBreathNodes(bst)

	for depth, nodes := range result{
		fmt.Printf("dept: %d\nnodes:\t", depth)
		for ; nodes.Len() > 0 ;{
			frontElement := nodes.Front()
			frontNode := frontElement.Value.(*BSTNode)
			fmt.Printf("%d\t", frontNode.Data)
			nodes.Remove(frontElement)
		}
		fmt.Printf("\n")
	}
}