package treesandgraphs

import "fmt"

func IsIncluded(bigTree BinaryTree, subTree BinaryTree) bool  {
	if bigTree.Root == nil && subTree.Root == nil{
		return true
	}
	if bigTree.Root == nil || subTree.Root == nil{
		return false
	}
	return IsEqual(subTree.Root, bigTree.Find(subTree.Root.Data))
}

func IsEqual(node1 *BinaryTreeNode, node2 *BinaryTreeNode) bool {
	if node1 == nil {
		return true
	}
	if node2 == nil || node1.Data != node2.Data {
		return false
	}

	return IsEqual(node1.LeftChild, node2.LeftChild) && IsEqual(node1.RightChild, node2.RightChild)
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

	fmt.Printf("\nfound:%+v", IsIncluded(tree, subtree))
}