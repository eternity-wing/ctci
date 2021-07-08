package treesandgraphs

func GetNextSuccessor(node *BinaryTreeNode) *BinaryTreeNode {
	if node == nil {
		return nil
	}
	if node.RightChild != nil {
		return GetMinNode(node.RightChild)
	}
	return GetGreaterAncestor(node)
}


