package treesandgraphs

import (
	"fmt"
)


func GetPathToTheNodeExample() {
	mainTree := BinaryTree{}
	leftTree := BinaryTree{}

	rightTree := BinaryTree{}
	rightTree.Insert(NewBinaryTreeNode(2))
	rightTree.Insert(NewBinaryTreeNode(5))
	rightTree.Insert(NewBinaryTreeNode(6))

	four := NewBinaryTreeNode(4)
	four.LeftChild = NewBinaryTreeNode(8)
	four.RightChild = NewBinaryTreeNode(9)
	leftTree.Insert(NewBinaryTreeNode(1))
	leftTree.Insert(four)
	leftTree.Insert(NewBinaryTreeNode(3))

	mainTree.Insert(NewBinaryTreeNode(0))
	mainTree.Insert(leftTree.Root)
	mainTree.Insert(rightTree.Root)

	ancestor := GetFirstCommonAncestor(mainTree.Root, mainTree.Root.LeftChild, four.RightChild)
	fmt.Printf("\nancestor:%+v\n", ancestor)
}

func GetFirstCommonAncestor(root *BinaryTreeNode, node1 *BinaryTreeNode, node2 *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	} else if root.Data == node1.Data || root.Data == node2.Data {
		return root
	}
	leftSubTree := GetFirstCommonAncestor(root.LeftChild, node1, node2)
	rightSubTree := GetFirstCommonAncestor(root.RightChild, node1, node2)

	if leftSubTree != nil && rightSubTree != nil {
		return root
	} else if leftSubTree != nil {
		return leftSubTree
	} else if rightSubTree != nil {
		return rightSubTree
	}
	return nil
}
