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

	ancestor := GetFirstCommonAncestor(mainTree.Root, four, four.RightChild)
	fmt.Printf("\nancestor:%+v\n", ancestor)
}

func GetFirstCommonAncestor(root *BinaryTreeNode, node1 *BinaryTreeNode, node2 *BinaryTreeNode) *BinaryTreeNode {
	if root == nil || root == node1 || root == node2 {
		return root
	}
	node1IsInLeft := Covers(root.LeftChild, node1)
	node2IsInLeft := Covers(root.LeftChild, node2)

	if node1IsInLeft != node2IsInLeft {
		return root
	}
	childSide := root.LeftChild
	if !node1IsInLeft {
		childSide = root.RightChild
	}
	return GetFirstCommonAncestor(childSide, node1, node2)
}

func Covers(root *BinaryTreeNode, node *BinaryTreeNode) bool {
	if root == nil {
		return false
	} else if root == node{
		return true
	}
	return Covers(root.LeftChild, node) || Covers(root.RightChild, node)
}


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func traverse(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	left, lca := traverse(root.Left, p, q)
	if left == 2 {
		return left, lca
	}

	center := 0
	if (root == p) || (root == q) {
		center = 1
	}

	if left+center == 2 {
		return left + center, root
	}

	right, lca := traverse(root.Left, p, q)
	if right == 2 {
		return right, lca
	}

	if left+center+right == 2 {
		return 2, root
	}

	return left + center + right, nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var lca *TreeNode
	var traverse func(root, p, q *TreeNode) bool
	traverse = func(root, p, q *TreeNode) bool {
		if root == nil {
			return false
		}

		if lca != nil {
			return false
		}

		left := 0
		if traverse(root.Left, p, q) {
			left = 1
		}

		right := 0
		if traverse(root.Right, p, q) {
			right = 1
		}

		center := 0
		if (root == p) || (root == q) {
			center = 1
		}

		if left+center+right == 2 {
			lca = root
		}

		return (left + center + right) > 0
	}

	traverse(root, p, q)

	return lca

}
