package treesandgraphs

import (
	"container/list"
	"fmt"
)

type BinaryTree struct {
	Root *BinaryTreeNode
}

func (b *BinaryTree) Insert(newNode *BinaryTreeNode)  {
	if newNode == nil{
		return
	}
	if b.Root == nil{
		b.Root = newNode
		return
	}

	traverseQueue := list.New()
	traverseQueue.PushBack(b.Root)

	for ; traverseQueue.Len() != 0;{
		nodeElement := traverseQueue.Front()
		node := nodeElement.Value.(*BinaryTreeNode)
		traverseQueue.Remove(nodeElement)
		if node.Data == newNode.Data{
			return
		}
		if node.LeftChild == nil {
			node.LeftChild = newNode
			newNode.Parent = node
			return
		}else{
			traverseQueue.PushBack(node.LeftChild)
		}

		if node.RightChild == nil {
			node.RightChild = newNode
			newNode.Parent = node
			return
		}else{
			traverseQueue.PushBack(node.RightChild)
		}

	}

}


func (b *BinaryTree) PrintInOrder()  {
	PrintInOrder(b.Root)
}


func (b *BinaryTree) Find(value int) *BinaryTreeNode {
	return b.findNodeByValue(b.Root, value)
}

func (b *BinaryTree) findNodeByValue(node *BinaryTreeNode, value int) (found *BinaryTreeNode ){
	if node == nil || node.Data == value{
		return node
	}

	if found = b.findNodeByValue(node.LeftChild, value); found != nil {
		return found
	}

	if found = b.findNodeByValue(node.RightChild, value); found != nil {
		return found
	}
	return nil
}

func RunExampleOfBinarySearchTree() {
	tree := BinaryTree{}
	tree.Insert(NewBinaryTreeNode(10))
	tree.Insert(NewBinaryTreeNode(15))
	tree.Insert(NewBinaryTreeNode(13))
	tree.Insert(NewBinaryTreeNode(16))
	tree.Insert(NewBinaryTreeNode(17))
	tree.Insert(NewBinaryTreeNode(5))
	tree.Insert(NewBinaryTreeNode(3))
	tree.Insert(NewBinaryTreeNode(4))
	tree.Insert(NewBinaryTreeNode(6))
	tree.PrintInOrder()
	found := tree.Find(10)
	fmt.Printf("\nfound:%+v", found)
}

