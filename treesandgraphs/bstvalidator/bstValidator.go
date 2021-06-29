package bstvalidator

import (
	"github.com/eternity-wings/dev_talks/treesandgraphs"
	"math"
)

func IsTreeValid(tree treesandgraphs.BinarySearchTree) bool  {
	return validateBSTNode(tree.Root, math.Inf(-1), math.Inf(1))
}


func validateBSTNode(node *treesandgraphs.BSTNode, min float64, max float64) bool  {
	if node == nil {
		return true
	}
	nodeData := float64(node.Data)
	if nodeData <= min || nodeData > max {
		return false
	}
	return validateBSTNode(node.LeftChild, min, nodeData) && validateBSTNode(node.RightChild, nodeData, max)

}