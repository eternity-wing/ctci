package treesandgraphs

import (
	"math/rand"
)

type BinarySearchTree struct {
	Root *BinaryTreeNode
}

func (b *BinarySearchTree) Add(value int) {
	newNode := &BinaryTreeNode{
		Data: value,
		Size: 1,
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
func (b *BinarySearchTree) GetRandomNode1() *BinaryTreeNode {
	return b.getRandomNode1(b.Root)
}

func (b *BinarySearchTree) getRandomNode(node *BinaryTreeNode) *BinaryTreeNode {
	leftSize := b.GetNodeSize(node.LeftChild)
	singleNodeProbability := float64(1) / float64(b.GetNodeSize(node))

	nodeProbability := singleNodeProbability
	leftSubTreeProbability := float64(leftSize) * singleNodeProbability


	roll := rand.Float64()*(1)

	// 0.5
	// [ 0.1, 0.5, 0.4 ]
	if roll <= nodeProbability {
		return node
	} else if roll <= nodeProbability+leftSubTreeProbability {
		return b.getRandomNode(node.LeftChild)
	} else {
		return b.getRandomNode(node.RightChild)
	}
}

func (b *BinarySearchTree) getRandomNode1(node *BinaryTreeNode) *BinaryTreeNode {
	index := rand.Intn(b.GetNodeSize(node))
	randomNode := node
	counter := 0
	for ; ; {
		if counter == index {
			return randomNode
		}
		leftSubtreeSize := b.GetNodeSize(randomNode.LeftChild)
		if index-counter > leftSubtreeSize {
			counter += leftSubtreeSize + 1
			randomNode = randomNode.RightChild
		} else {
			counter += 1
			randomNode = randomNode.LeftChild
		}
	}
}

func (b *BinarySearchTree) GetNodeSize(node *BinaryTreeNode) int {
	if node == nil {
		return 0
	}
	return node.Size
}
