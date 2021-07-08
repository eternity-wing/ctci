package treesandgraphs

import "container/list"

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
}