package treesandgraphs

import "math"

func IsBalance(node *BinaryTreeNode, height *int) bool {
	if node == nil {
		*height = 0
		return true
	}
	leftHeight := new(int)
	rightHeight := new(int)
	isLeftBalance := IsBalance(node.LeftChild, leftHeight)
	isRightBalance := IsBalance(node.RightChild, rightHeight)
	*height = int(math.Max(float64(*leftHeight), float64(*rightHeight)) + 1)
	balanceFactor := *leftHeight - *rightHeight
	return (balanceFactor <= 1 && balanceFactor >= -1) && isLeftBalance && isRightBalance

}


// O(size(node))
func GetHeight(root *BinaryTreeNode) int {
	if root == nil {
		return -1
	}
	return int(math.Max(float64(GetHeight(root.LeftChild)), float64(GetHeight(root.RightChild))) + 1)

}
