package treesandgraphs

import (
	"math/rand"
	"time"
)

type BinarySearchTree struct {
	Root *BinaryTreeNode
}

func (b *BinarySearchTree) Add(value int) {
	newNode := &BinaryTreeNode{
		Data:                  value,
		NumberOfGrandChildren: 1,
	}

	if b.Root == nil {
		b.Root = newNode
		return
	}

	for node := b.Root; node != nil; {
		if node.Data == newNode.Data {
			return
		}
		if node.Data < newNode.Data && node.RightChild == nil {
			node.RightChild = newNode
			newNode.Parent = node
			node.IncrementAncestorsNumberOfGrandChildren()
			return
		} else if node.Data < newNode.Data {
			node = node.RightChild
			continue
		}

		if node.Data > newNode.Data && node.LeftChild == nil {
			node.LeftChild = newNode
			newNode.Parent = node
			node.IncrementAncestorsNumberOfGrandChildren()
			return
		} else if node.Data > newNode.Data {
			node = node.LeftChild
			continue
		}

	}

}

func (b *BinarySearchTree) IsBalance() bool {
	if b.Root == nil {
		return true
	}
	balanceFactor := GetBSTNodeBalanceFactor(*b.Root)
	return balanceFactor <= 1 && balanceFactor >= -1
}

func (b *BinarySearchTree) PrintInOrder() {
	PrintInOrder(b.Root)
}

func (b *BinarySearchTree) GetRandomNode() *BinaryTreeNode {
	return b.getRandomNode(b.Root)
}

func (b *BinarySearchTree) getRandomNode(node *BinaryTreeNode) *BinaryTreeNode {
	numberOfLeftChildGrandChildren := b.GetNumberOfGrandChildren(node.LeftChild)
	singleNodeProbability := float64(1) / float64(b.GetNumberOfGrandChildren(node))

	nodeProbability := 1 * singleNodeProbability
	leftSubTreeProbability := float64(numberOfLeftChildGrandChildren) * singleNodeProbability

	rand.Seed(time.Now().UnixNano())
	roll := 0 + rand.Float64()*(1)

	if roll <= nodeProbability {
		return node
	} else if roll <= nodeProbability+leftSubTreeProbability {
		return b.getRandomNode(node.LeftChild)
	} else {
		return b.getRandomNode(node.RightChild)
	}
}

func (b *BinarySearchTree) GetNumberOfGrandChildren(node *BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return node.NumberOfGrandChildren
}
