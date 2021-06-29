package avltrees

import "math"

func GetAvlNodeHeight(node *AVLNode) int {
	if node == nil {
		return -1
	}
	return node.Height
}

func CalculateNodeHeight(node *AVLNode) {
	if node == nil {
		return
	}
	node.Height = int(math.Max(float64(GetAvlNodeHeight(node.LeftChild)), float64(GetAvlNodeHeight(node.RightChild)))) + 1
}