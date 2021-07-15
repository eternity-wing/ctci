package treesandgraphs

import "fmt"

func IsSubTree(tree BinaryTree, subTree BinaryTree) bool  {
	if tree.Root == nil && subTree.Root == nil{
		return true
	}
	if tree.Root == nil || subTree.Root == nil{
		return false
	}
	return IsSubNode(subTree.Root, tree.Find(subTree.Root.Data))
}

func IsSubNode(node1 *BinaryTreeNode, node2 *BinaryTreeNode) bool {
	if node1 == nil {
		return true
	}
	if node2 == nil || node1.Data != node2.Data {
		return false
	}

	return IsSubNode(node1.LeftChild, node2.LeftChild) && IsSubNode(node1.RightChild, node2.RightChild)
}



func RunExampleOfIsSubtree() {
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

	subtree := BinaryTree{}
	node16 := NewBinaryTreeNode(16)
	node16.RightChild = NewBinaryTreeNode(4)
	subtree.Insert(node16)

	fmt.Printf("\nfound:%+v", IsSubTree(tree, subtree))
}